.PHONY: all hello

all: hello

hello:
	@printf "Hello lovely human\n"

#---------------
#-- Local dev
#---------------
.PHONY: dev
dev: deps

deps:
	@printf "Downloading go deps...\n"
	@go mod tidy && go mod vendor
	@printf "Downloaded go deps...\n"

tests:
	@printf "Running tests...\n"
	@go test ./... -count=1 --tags=integration_test -race
	@printf "Run tests.\n"

docker:
	@printf "Building container...\n"
	@docker build . -f ./build/package/cli.Dockerfile -t grpc-stub
	@printf "built container.\n"

#---------------
#-- Code linting
#---------------
.PHONY: lint
lint:
	make download-lint
	@echo "Linting code...\n"
	@golangci-lint run --fix
	@echo "Linted code.\n"

check-lint:
	make download-lint
	@echo "Checking lint code...\n"
	@golangci-lint run
	@echo "Checked lint code.\n"

download-lint:
	@printf "Downloading golangci-lint...\n"
	go install -mod=readonly github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
	@printf "Downloaded golangci-lint.\n"
