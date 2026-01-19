NAME           := ipdrawer
RELEASE_TAG    ?= v2.1
VERSION        := $(shell git describe --tags --exact-match 2> /dev/null || git rev-parse --short HEAD || echo "unknown")
REVISION       := $(shell git rev-parse HEAD)
SRCS           := $(shell find . -type f -name '*.go')
PROTOSRCS      := $(shell find . -type f -name '*.proto' | grep -v -e vendor | grep -v -e node_modules)
LINUX_LDFLAGS  := -s -w -extldflags "-static"
DARWIN_LDFLAGS := -s -w
LINKFLAGS      := \
	-X "github.com/hatena/ipdrawer/pkg/build.tag=$(VERSION)" \
	-X "github.com/hatena/ipdrawer/pkg/build.rev=$(REVISION)"
override LINUX_LDFLAGS += $(LINKFLAGS)
override DARWIN_LDFLAGS += $(LINKFLAGS)

PKG                  := github.com/hatena/ipdrawer
API_CLIENT_DIR       := gen/client
API_SPEC             := gen/openapiv2/serverpb/server.swagger.json
SWAGGER_UI_DATA_PATH := pkg/ui/
SWAGGER_UI_SRC       := third_party/swagger-ui
DOCKER_PROXY_ARGS    := --build-arg GO_PROXY="https://proxy.golang.org,direct"

$(NAME): $(SRCS)
	go build -ldflags '$(DARWIN_LDFLAGS)' $(PKG)/cmd/ipdrawer

.PHONY: cross-build
cross-build:
	GOOS=darwin GOARCH=amd64 go build -ldflags '$(DARWIN_LDFLAGS)' -o dist/$(NAME)_darwin_amd64 $(PKG)/cmd/...
	GOOS=linux GOARCH=amd64 go build -a -tags netgo -installsuffix netgo -ldflags '$(LINUX_LDFLAGS)' -o dist/$(NAME)_linux_amd64 $(PKG)/cmd/...

.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 go build -a -tags netgo -installsuffix netgo -ldflags '$(LINUX_LDFLAGS)' -o dist/$(NAME)_linux_amd64 $(PKG)/cmd/...

.PHONY: vet
vet:
	go vet -all -printfuncs=Wrap,Wrapf,Errorf ./...

.PHONY: test
test:
	go test -cover $$(go list ./... | grep -v -e node_modules)

.PHONY: test-race
test-race:
	go test -v -race $$(go list ./... | grep -v -e node_modules)

.PHONY: test-all
test-all: vet test-race

.PHONY: fmt
fmt:
	gofmt -s -w $$(find . -type f -name '*.go' | grep -v -e vendor -e node_modules)

.PHONY: imports
imports:
	goimports -w $$(find . -type f -name '*.go' | grep -v -e vendor -e node_modules)

# Proto is generated and saved in gen/go.
# Re-make this target if proto file changes.
.PHONY: proto
proto:
	rm -rf gen
	buf generate proto
	go generate ./tools
	go install golang.org/x/tools/cmd/goimports@latest
	make gen-client-docker
	make fmt imports
	go mod tidy

.PHONY: ui
ui:
	statik -dest $(SWAGGER_UI_DATA_PATH) -p swagger -src $(SWAGGER_UI_SRC)
	make fmt imports

.PHONY: gen-client-docker
gen-client-docker: $(API_SPEC)
	docker run --rm -v ${PWD}:/local cylonix/openapi-generator-cli:v7.8.5 \
		generate -g go \
		-i /local/$(API_SPEC) \
		-o /local/$(API_CLIENT_DIR) \
		--git-repo-id ipdrawer/$(API_CLIENT_DIR) --git-user-id hatena \
		--additional-properties packageName=apiclient,enumClassPrefix=true,packageVersion=1.0 \
		--inline-schema-options RESOLVE_INLINE_ENUMS=true
	sudo chown -R ${USER}:${USER} ./$(API_CLIENT_DIR)
	@rm -rf \
		$(API_CLIENT_DIR)/go.mod \
		$(API_CLIENT_DIR)/go.sum

.PHONY: docker
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
docker:
	docker build ${DOCKER_PROXY_ARGS} \
		--network host \
		--build-arg VERSION=$(VERSION) \
		--build-arg VCS_URL="https://gitlab.com/cylonix/sase/ipdrawer" \
		--build-arg VCS_REF=$(REVISION) \
		--build-arg VCS_BRANCH=$(BRANCH) \
		--build-arg BUILD_DATE=$(BUILD_DATE) \
		--tag cylonix/ipdrawer:$(REVISION) \
		--tag cylonix/ipdrawer:$(VERSION) \
		--tag cylonix/ipdrawer:$(RELEASE_TAG) \
		--tag cylonix/ipdrawer:latest \
		.

.PHONY: clean
clean:
	rm -rf $(NAME) dist

# Stress tests for IPAM performance
.PHONY: stress-test
stress-test:
	@echo "Running IPAM stress tests (/16 and /18 pools)..."
	go test ./pkg/ipam/... -v -run "TestStressLargePool|TestStressCompareAlgorithms" -timeout 10m

.PHONY: stress-test-large
stress-test-large:
	@echo "Running IPAM stress test on /12 pool (1M IPs)..."
	@echo "This test requires RUN_VERY_LARGE_POOL_TEST=1 environment variable"
	RUN_VERY_LARGE_POOL_TEST=1 go test ./pkg/ipam/... -v -run "TestStressVeryLargePool" -timeout 20m

.PHONY: stress-test-production
stress-test-production:
	@echo "Running IPAM stress test on /10 pool (4M IPs) - Production scale..."
	@echo "This test requires RUN_PRODUCTION_POOL_TEST=1 environment variable"
	@echo "WARNING: This test takes ~3 minutes and uses significant memory"
	RUN_PRODUCTION_POOL_TEST=1 go test ./pkg/ipam/... -v -run "TestStressProductionPool" -timeout 30m

.PHONY: stress-test-all
stress-test-all: stress-test stress-test-large stress-test-production
	@echo "All stress tests completed"
