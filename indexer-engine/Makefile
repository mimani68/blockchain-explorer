.PHONY: all
all: build
FORCE: ;

SHELL  := env BOOKMARK_ENV=$(BOOKMARK_ENV) $(SHELL)
BOOKMARK_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf ./build/*

dependencies:
	go mod download

build: dependencies clean build-api # build-cmd

build-api: 
	go build -tags $(BOOKMARK_ENV) -o ./build/api_server init/main.go

# build-cmd:
# 	go build -tags $(BOOKMARK_ENV) -o ./build/search cmd/main.go

linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(BOOKMARK_ENV) netgo" -installsuffix netgo -o $(BIN_DIR)/api api/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(BOOKMARK_ENV) netgo" -installsuffix netgo -o $(BIN_DIR)/search cmd/main.go

ci: dependencies test	

grpc_builder:
	@protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--swagger_out=logtostderr=true:. \
		internal/transport/grpc/service/*.proto
	@echo done

gosec:
	@echo "Run everything except for rule G303"
	@gosec -exclude=G303 ./...
	@echo "https://github.com/securego/gosec"

run_root:
	@go run init/main.go

run_grpc:
	@go run init/main.go grpc --port 3000 --host 0.0.0.0 --env development

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=pkg/bookmark/interface.go -destination=pkg/bookmark/mock/bookmark.go -package=mock

test:
	@echo unit test whole part of system
	go test -tags testing ./...

install_swagger:
	@download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
  		jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
	@sudo curl -o /usr/local/bin/swagger -L'#' "$download_url"
	@sudo chmod +x /usr/local/bin/swagger

swagger_doc_serve: swag_doc_generator
	@swagger serve --flatten --port=4000 --path=/api/docs/swagger ./docs/swagger.json

swagger_docker_ui: swag_doc_generator
	docker run --rm --name swg-ui -p 3080:8080 -e SWAGGER_JSON=/files/swagger.json -v ${PWD}/docs/swagger:/files swaggerapi/swagger-ui:v4.12.0

swag_doc_generator:
	@echo https://github.com/swaggo/echo-swagger
	@swag init -g init/main.go --output=docs/swagger

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
