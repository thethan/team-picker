PROJECT_NAME := "team_picker"
PKG := "github.com/thethan/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
GOARCH=wasm
GOOS=js
PORT=8080

wasm: ## Build the wasm file for packages
    GOOS=js GOARCH=wasm go build -o lib.wasm lib.go

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Vet the files
	@go vet ${PKG_LIST}

test: ## Run unit tests
	@go test -short ${PKG_LIST}

testall: ## Run unit tests
	@go test ${PKG_LIST}

dep: ## Update dependencies
	@dep ensure -vendor-only

race: ## Run unit tests
	@go test -race -short ${PKG_LIST}

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the release and develoment container. The development
	docker build -t $(PROJECT_NAME) .

run: ## Run container on port configured in `config.env`
	docker run -p=$(PORT):$(PORT) --name=$(PROJECT_NAME) $(PROJECT_NAME)

# The `validate` target checks for errors and inconsistencies in
# our specification of an API. This target can check if we're
# referencing inexistent definitions and gives us hints to where
# to fix problems with our API in a static manner.
validate:
	swagger validate ./swagger/swagger.yml

# The `gen` target depends on the `validate` target as
# it will only succesfully generate the code if the specification
# is valid.
#
# Here we're specifying some flags:
# --target              the base directory for generating the files;
# --spec                path to the swagger specification;
# --exclude-main        generates only the library code and not a
#                       sample CLI application;
# --name                the name of the application.
gen: validate
	swagger generate server \
		--target=./swagger \
		--spec=./swagger/swagger.yml \
		--exclude-main \
		--name=hello
