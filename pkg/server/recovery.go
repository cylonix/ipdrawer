package server

import (
	"runtime"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	MAXSTACKSIZE = 4096
)

func recoveryFunc(p interface{}) error {
	stack := make([]byte, MAXSTACKSIZE)
	stack = stack[:runtime.Stack(stack, false)]
	logger.Errorf("panic grpc: err=%v, stack:\n%s", p, string(stack))
	return status.Errorf(codes.Internal, "panic error: %v", p)
}
