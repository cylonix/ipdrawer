package server

import (
	ocontext "context"
	"errors"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/opentracing/opentracing-go"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"github.com/urfave/negroni"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/hatena/ipdrawer/gen/go/serverpb"
	"github.com/hatena/ipdrawer/pkg/base"
	"github.com/hatena/ipdrawer/pkg/ipam"
	_ "github.com/hatena/ipdrawer/pkg/ui/swagger" // import static files
)

var logrusEntry = logrus.NewEntry(logrus.New())

type APIServer struct {
	serverpb.UnimplementedNetworkServiceV0Server
	serverpb.UnimplementedIPServiceV0Server
	serverpb.UnimplementedPoolServiceV0Server
	lis     net.Listener
	manager *ipam.IPManager
	grpcS   *grpc.Server
	httpS   *http.Server
}

func NewServer(cfg *base.Config) *APIServer {
	mngr := ipam.NewIPManager(cfg)
	lis, err := net.Listen("tcp", ":"+cfg.Port)

	if err != nil {
		panic(err)
	}
	return &APIServer{
		lis:     lis,
		manager: mngr,
	}
}

func (api *APIServer) newGateway(ctx context.Context) (http.Handler, error) {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers:  true,
				EmitUnpopulated: true,
				Indent:          "  ",
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)
	addr := api.lis.Addr().String()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := serverpb.RegisterNetworkServiceV0HandlerFromEndpoint(
		ctx, mux, addr, opts); err != nil {
		return nil, err
	}
	if err := serverpb.RegisterIPServiceV0HandlerFromEndpoint(
		ctx, mux, addr, opts); err != nil {
		return nil, err
	}
	if err := serverpb.RegisterPoolServiceV0HandlerFromEndpoint(
		ctx, mux, addr, opts); err != nil {
		return nil, err
	}

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negronilogrus.NewMiddleware())
	n.UseHandler(mux)

	return mux, nil
}

func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, http.FileServer(statikFS)))

	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(serverpb.SWAGGER))
	})
}

func (api *APIServer) Start() error {
	log := logger.WithField("handle", "Server")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cm := cmux.New(api.lis)
	grpcL := cm.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := cm.Match(cmux.HTTP1Fast())

	api.grpcS = grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(
				grpc_recovery.WithRecoveryHandler(recoveryFunc)),
			grpc_logrus.StreamServerInterceptor(logrusEntry),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(
				grpc_opentracing.WithTracer(opentracing.GlobalTracer()),
			),
			grpc_recovery.UnaryServerInterceptor(
				grpc_recovery.WithRecoveryHandler(recoveryFunc)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		)),
	)

	// Register gRPC server to api
	serverpb.RegisterNetworkServiceV0Server(api.grpcS, api)
	serverpb.RegisterIPServiceV0Server(api.grpcS, api)
	serverpb.RegisterPoolServiceV0Server(api.grpcS, api)

	gw, err := api.newGateway(ctx)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", gw)
	serveSwagger(mux)
	api.httpS = &http.Server{
		Handler: mux,
	}

	go func() {
		if err := api.grpcS.Serve(grpcL); err != nil {
			log.WithError(err).Errorln("grpc server exited")
			if !errors.Is(err, cmux.ErrServerClosed) {
				panic(err)
			}
		}
	}()
	go func() {
		if err := api.httpS.Serve(httpL); err != nil {
			log.WithError(err).Errorln("http server exited")
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	return cm.Serve()
}

func (api *APIServer) Shutdown(
	ctx ocontext.Context,
	stopped chan struct{},
) {
	api.grpcS.GracefulStop()
	api.httpS.Shutdown(ctx)
	stopped <- struct{}{}
}

func ParseIPAndMask(ip, mask string) (net.IP, error) {
	return nil, nil
}

func init() {
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)
}
