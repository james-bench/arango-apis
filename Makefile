SCRIPTDIR := $(shell pwd)
ROOTDIR := $(shell cd $(SCRIPTDIR) && pwd)
BUILDIMAGE := arangodb-cloud-apis-build
CACHEVOL := arangodb-cloud-apis-gocache

DOCKERARGS := run -t --rm \
	-u $(shell id -u):$(shell id -g) \
	-v $(ROOTDIR)/vendor:/go/src \
	-v $(ROOTDIR):/usr/src \
	-v $(CACHEVOL):/usr/gocache \
	-e GOCACHE=/usr/gocache \
	-e CGO_ENABLED=0 \
	-w /usr/src \
	$(BUILDIMAGE)

.PHONY: all
all: generate

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

# Generate go code for proto files
.PHONY: generate
generate: $(CACHEVOL) 
	docker $(DOCKERARGS) \
		go generate ./...
