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
all: generate docs

# Build docker builder image
.PHONY: build-image
build-image:
ifndef CIRCLECI
	docker build -t $(BUILDIMAGE) -f Dockerfile.build .
else
	go get github.com/gogo/protobuf/protoc-gen-gogo
	go get github.com/gogo/protobuf/protoc-gen-gofast
	go get github.com/gogo/protobuf/protoc-gen-gogofaster
	go get golang.org/x/tools/cmd/goimports
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
endif

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

# Generate API docs
.PHONY: docs
docs: $(CACHEVOL) $(MODVOL)
	$(DOCKERENV) \
		protoc -I.:vendor:vendor/googleapis/:vendor/github.com/gogo/protobuf/protobuf/ \
			--doc_out=docs $(PROTOSOURCES) \
			--doc_opt=markdown,apis.md

.PHONY: test
test:
	mkdir -p bin/test
	go test -coverprofile=bin/test/coverage.out -v ./... | tee bin/test/test-output.txt ; exit "$${PIPESTATUS[0]}"
	cat bin/test/test-output.txt | go-junit-report > bin/test/unit-tests.xml
	go tool cover -html=bin/test/coverage.out -o bin/test/coverage.html

bootstrap:
	go get github.com/jstemmer/go-junit-report
