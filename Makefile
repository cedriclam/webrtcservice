# Makefile

PROJECTNAME := "webrtcservice"
ROOT := $(PWD)
BINFOLDER := $(ROOT)/bin

# GO WORKSPACE
GO := go
GOPATH := $(ROOT)
WORKSPACE := $(ROOT)/src/$(PROJECTNAME)

# Hack tools
GLIDEPATH := $(PWD)/hack/glide
GLIDE := $(GLIDEPATH)/bin/glide
GOLINTGOPATH := $(PWD)/hack/golint
GOLINT := $(GOLINTGOPATH)/bin/golint
GO2XUNITGOPATH := $(PWD)/hack/go2xunit
GO2XUNIT := $(GO2XUNITGOPATH)/bin/go2xunit
GOCOVGOPATH := $(PWD)/hack/gocov
GOCOV := $(GOCOVGOPATH)/bin/gocov
GOCOVXMLGOPATH := $(PWD)/hack/gocov-xml
GOCOVXML := $(GOCOVXMLGOPATH)/bin/gocov-xml


# DOCKER Folder
DOCKERFOLDER := $(ROOT)/docker

# These are the values we want to pass for Version and BuildTime
VERSION := $(shell git describe --tags --always --dirty)
TAG ?= $(VERSION)
BUILD_TIME=$(shell date +%FT%T%z)

# Common env var
COMMONENVVAR := GO15VENDOREXPERIMENT=1 GOPATH=$(GOPATH)
COMPONENTS :=  $(shell cd $(WORKSPACE) && $(GLIDE) nv)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X $(PROJECTNAME)/main.Version=${VERSION} -X $(PROJECTNAME)/main.BuildTime=${BUILD_TIME}"


# Generic targets
.PHONY: all
all: $(GLIDE) $(GEN) $(GOLINT) $(GO2XUNIT) $(GOCOV) $(GOCOVXML) vendor build validate test docker ## Install tools and run the following targets: vendor, build, validate, test, docker

.PHONY: init
init: get-tools vendor ## Init project

.PHONY: vendor
vendor: $(GLIDE) ## Install all vendor dependencies
	@cd $(WORKSPACE) && $(GLIDE) install

.PHONY: docker
docker: build-docker ## Build docker images
	$(MAKE) -C docker build

.PHONY: validate
validate: $(GOLINT) ## Run go style validation (golint)
	@hack/check-go-style

reports:
	@mkdir -p _reports

.PHONY: cover
cover: ## Run all test coverage
	@cd $(WORKSPACE) && $(COMMONENVVAR) $(GO) test $(COMPONENTS) --cover

.PHONY: cover-extra
cover-extra: $(GLIDE) $(GOCOV) ## Run test coverage and generate report in standard output
	@cd $(WORKSPACE) && $(COMMONENVVAR) $(GOCOV) test $(COMPONENTS) | $(GOCOV) report

.PHONY: cover-xml
cover-xml: reports  $(GLIDE) $(GOCOV) $(GOCOVXML) ## Run test coverage and generate report in _reports folder
	@cd $(WORKSPACE) && $(COMMONENVVAR) $(GOCOV) test $(COMPONENTS) | $(GOCOVXML) > $(ROOT)/_reports/cover_all.xml

.PHONY: test-xml
test-xml: reports  $(GLIDE) $(GO2XUNIT) ## Run all tests and generate reports in _reports folder
	@cd $(WORKSPACE) && $(COMMONENVVAR) $(GO) test -v $(COMPONENTS) | $(GO2XUNIT) --output $(ROOT)/_reports/test_all.xml

.PHONY: build
build: ## Build rest server binary
	@cd $(WORKSPACE)/ &&\
	$(COMMONENVVAR) $(GO) build \
	-o $(BINFOLDER)/$(PROJECTNAME) \
	$(LDFLAGS)

.PHONY: build-docker
build-docker: ## Build rest server docker image
	@echo "Building redismanager docker image..."
	@cd $(WORKSPACE)/ &&\
	$(COMMONENVVAR) GOOS=linux GOARCH=amd64 CGO_ENABLED=1 $(GO) build \
	-o $(DOCKERFOLDER)/$(PROJECTNAME)/$(PROJECTNAME) \
	$(LDFLAGS)

.PHONY: test
test: ## Run unit-tests
	@cd $(WORKSPACE)/ && $(COMMONENVVAR) $(GO) test $(COMPONENTS)

.PHONY: run
run: ## Run currencyconverter server process
	@cd $(WORKSPACE)/ && $(COMMONENVVAR) $(GO) run main.go


# Hack targets build
.PHONY: get-tools
get-tools: ## Download all tools dependencies in hack sub-directory
	@mkdir -p $(GLIDEPATH) && cd $(GLIDEPATH) && GOPATH=$(GLIDEPATH) go get github.com/Masterminds/glide
	@mkdir -p $(GOLINTGOPATH) && cd $(GOLINTGOPATH) && GOPATH=$(GOLINTGOPATH) go get github.com/golang/lint/golint
	@mkdir -p $(GO2XUNITGOPATH) && cd $(GO2XUNITGOPATH) && GOPATH=$(GO2XUNITGOPATH) go get bitbucket.org/tebeka/go2xunit
	@mkdir -p $(GOCOVGOPATH) && cd $(GOCOVGOPATH) && GOPATH=$(GOCOVGOPATH) go get github.com/axw/gocov/gocov
	@mkdir -p $(GOCOVXMLGOPATH) && cd $(GOCOVXMLGOPATH) && GOPATH=$(GOCOVXMLGOPATH) go get github.com/AlekSi/gocov-xml

$(GLIDE):
	@cd $(GLIDEPATH) && GO15VENDOREXPERIMENT=1 GOPATH=$(GLIDEPATH) go get github.com/Masterminds/glide

$(GOLINT):
	@cd $(GOLINTGOPATH) && GOPATH=$(GOLINTGOPATH) go install github.com/golang/lint/golint

$(GO2XUNIT):
	@cd $(GO2XUNITGOPATH) && GOPATH=$(GO2XUNITGOPATH) go install bitbucket.org/tebeka/go2xunit

$(GOCOV):
	@cd $(GOCOVGOPATH) && GOPATH=$(GOCOVGOPATH) go install github.com/axw/gocov/gocov

$(GOCOVXML):
	@cd $(GOCOVXMLGOPATH) && GOPATH=$(GOCOVXMLGOPATH) go install github.com/AlekSi/gocov-xml

.PHONY: clean
clean: ## Clean project
	@rm -rf hack/glide/{bin,pkg}
	@rm -rf hack/golint/{bin,pkg}
	@rm -rf hack/go2xunit/{bin,pkg}
	@rm -rf hack/gocov/{bin,pkg}
	@rm -rf hack/gocov-xml/{bin,pkg}
	$(MAKE) -C docker clean
#> /dev/null 2>&1

.PHONY: help
help: ## Display list of targets
	@echo ''
	@echo '  Usage:'
	@echo '    make <target>'
	@echo '  Targets: '
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-22s\033[0m %s\n", $$1, $$2}'
	@echo ''

.DEFAULT_GOAL := help
