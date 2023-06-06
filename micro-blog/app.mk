GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif

CURRENT_DIR=$(shell a=`pwd` && echo $$a/)
API_PROTO=../api/proto/
API_PROTO_SED=$(shell echo $(API_PROTO) | sed 's/\//\\\//g')


# initialize develop environment
init:
	@which micro  &> /dev/null || go install github.com/devexps/go-micro/cmd/micro/v2@latest
	@which protoc-gen-go-http  &> /dev/null || go install github.com/devexps/go-micro/cmd/protoc-gen-go-http/v2@latest
	@which protoc-gen-go-errors  &> /dev/null || go install github.com/devexps/go-micro/cmd/protoc-gen-go-errors/v2@latest
	@which protoc-gen-go &> /dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
	@which protoc-gen-go-grpc &> /dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	@which protoc-gen-validate  &> /dev/null || go install github.com/envoyproxy/protoc-gen-validate@v1.0.0
	@which protoc-gen-openapi  &> /dev/null || go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@which buf  &> /dev/null || go install github.com/bufbuild/buf/cmd/buf@latest
	@which ent  &> /dev/null || go install entgo.io/ent/cmd/ent@latest
	@which golangci-lint  &> /dev/null || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.0

# generate protobuf api go code
proto:
	@cd $(CURRENT_DIR)$(API_PROTO) && buf generate

# generate OpenAPI v3 doc
openapi:
	@$(foreach dir, $(dir $(foreach dir, $(API_PROTO), $(wildcard $(dir)/*/*/buf.openapi.gen.yaml))),\
		$(eval DIR=$(shell echo .$(dir) | sed "s/$(API_PROTO_SED)//")) \
		cd $(CURRENT_DIR)$(API_PROTO) && buf generate --path $(DIR) --template $(DIR)buf.openapi.gen.yaml; \
	)

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo " make [target]"
	@echo ""
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help