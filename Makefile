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

.PHONY: all
all: generate docs

# Build docker builder image
.PHONY: build-image
build-image:
	docker build -t $(BUILDIMAGE) -f Dockerfile.build .

.PHONY: $(CACHEVOL)
$(CACHEVOL):
	@docker volume create $(CACHEVOL)
	docker run -it 	--rm -v $(CACHEVOL):/usr/gocache \
		$(BUILDIMAGE) \
		chown -R $(shell id -u):$(shell id -g) /usr/gocache

.PHONY: $(MODVOL)
$(MODVOL):
	@docker volume create $(MODVOL)
	docker run -it 	--rm -v $(MODVOL):/go/pkg/mod \
		$(BUILDIMAGE) \
		chown -R $(shell id -u):$(shell id -g) /go/pkg/mod

# Generate go code for proto files
.PHONY: generate
generate: $(CACHEVOL) $(MODVOL)
	docker $(DOCKERARGS) \
		go generate ./...

# Generate API docs
.PHONY: docs
docs: $(CACHEVOL) $(MODVOL)
	@echo $(PROTOSOURCES)
	docker $(DOCKERARGS) \
		protoc -I.:vendor:vendor/googleapis/:vendor/github.com/gogo/protobuf/protobuf/ \
			--doc_out=docs $(PROTOSOURCES) \
			--doc_opt=markdown,apis.md

.PHONY: test
test:
	go test ./...
