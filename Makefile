SHELL = bash
SCRIPTDIR := $(shell pwd)
ROOTDIR := $(shell cd $(SCRIPTDIR) && pwd)
BUILDIMAGE := arangodb-cloud-apis-build
CACHEVOL := arangodb-cloud-apis-gocache
MODVOL := arangodb-cloud-apis-pkg-mod
PROTOSOURCES := $(shell find .  -name '*.proto' -not -path './vendor/*')

DOCKERARGS := run -t --rm \
	-u $(shell id -u):$(shell id -g) \
	-v $(ROOTDIR)/vendor:/go/src \
	-v $(ROOTDIR):/usr/src \
	-v $(CACHEVOL):/usr/gocache \
	-v $(MODVOL):/go/pkg/mod \
	-e GOCACHE=/usr/gocache \
	-e CGO_ENABLED=0 \
	-w /usr/src \
	$(BUILDIMAGE)

ifndef CIRCLECI
	DOCKERENV := docker $(DOCKERARGS)
else
	DOCKERENV :=
endif

.PHONY: all
all: generate build check ts docs

# Build docker builder image
.PHONY: build-image
ifndef CIRCLECI
build-image:
	docker build \
		--build-arg=TOKEN=$(shell cat $(HOME)/.arangodb/ms/github-readonly-code-acces.token) \
		-t $(BUILDIMAGE) \
		-f Dockerfile.build .
else
build-image: get-plugins
endif

.PHONY: get-plugins
get-plugins:
	go get github.com/gogo/protobuf/protoc-gen-gogo@v1.3.0
	go get github.com/gogo/protobuf/protoc-gen-gofast@v1.3.0
	go get github.com/gogo/protobuf/protoc-gen-gogofaster@v1.3.0
	go get golang.org/x/tools/cmd/goimports@v0.0.0-20190918214516-5a1a30219888
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.11.0
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.11.0
	go get github.com/golang/protobuf/protoc-gen-go@v1.3.2
	go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.3.2
	go get github.com/arangodb-managed/protoc-gen-ts/cmd/protoc-gen-ts@v0.1.0

.PHONY: $(CACHEVOL)
$(CACHEVOL):
ifndef CIRCLECI
	@docker volume create $(CACHEVOL)
	docker run -it 	--rm -v $(CACHEVOL):/usr/gocache \
		$(BUILDIMAGE) \
		chown -R $(shell id -u):$(shell id -g) /usr/gocache
endif

.PHONY: $(MODVOL)
$(MODVOL):
ifndef CIRCLECI
	@docker volume create $(MODVOL)
	docker run -it 	--rm -v $(MODVOL):/go/pkg/mod \
		$(BUILDIMAGE) \
		chown -R $(shell id -u):$(shell id -g) /go/pkg/mod
endif

# Generate go code for proto files
.PHONY: generate
generate: $(CACHEVOL) $(MODVOL)
	$(DOCKERENV) \
		go generate ./...

# Build go code 
.PHONY: build
build: generate
	go build ./...

# Check go code 
.PHONY: check
check: 
	zutano go check ./...

# Generate API docs
.PHONY: docs
docs: $(CACHEVOL) $(MODVOL)
	$(DOCKERENV) \
		protoc -I.:vendor:vendor/googleapis/:vendor/github.com/gogo/protobuf/protobuf/ \
			--doc_out=docs $(PROTOSOURCES) \
			--doc_opt=html,index-raw.html
	cat docs/header.txt docs/index-raw.html > docs/index.html
	rm docs/index-raw.html

# Generate API as typescript
.PHONY: ts
ts: $(CACHEVOL) $(MODVOL)
	@rm -Rf typescript
	@mkdir -p typescript
	$(DOCKERENV) \
		protoc -I.:vendor:vendor/googleapis/:vendor/github.com/gogo/protobuf/protobuf/ \
			--ts_out=typescript $(PROTOSOURCES) \
			--ts_opt=.

.PHONY: test
test:
	mkdir -p bin/test
	go test -coverprofile=bin/test/coverage.out -v ./... | tee bin/test/test-output.txt ; exit "$${PIPESTATUS[0]}"
	cat bin/test/test-output.txt | go-junit-report > bin/test/unit-tests.xml
	go tool cover -html=bin/test/coverage.out -o bin/test/coverage.html

bootstrap:
	go get github.com/arangodb-managed/zutano
	go get github.com/jstemmer/go-junit-report
