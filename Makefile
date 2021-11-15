SHELL = bash
SCRIPTDIR := $(shell pwd)
ROOTDIR := $(shell cd $(SCRIPTDIR) && pwd)
BUILDIMAGE := arangodboasis/golang-ci:latest
CACHEVOL := arangodb-cloud-apis-gocache
MODVOL := arangodb-cloud-apis-pkg-mod
HOMEVOL := arangodb-cloud-apis-home
PROTOSOURCES := $(shell find .  -name '*.proto' -not -path './vendor/*' | sort)

ifndef CIRCLECI
	GITHUB_TOKEN := $(shell cat $(HOME)/.arangodb/ms/github-readonly-code-acces.token)
else
	GITHUB_TOKEN :=
endif

DOCKERARGS := run -t --rm \
	-u $(shell id -u):$(shell id -g) \
	-v $(ROOTDIR)/vendor:/go/src \
	-v $(ROOTDIR):/usr/src \
	-v $(CACHEVOL):/usr/gocache \
	-v $(MODVOL):/go/pkg/mod \
	-v $(HOMEVOL):/home/gopher \
	-e GOCACHE=/usr/gocache \
	-e GOSUMDB=off \
	-e CGO_ENABLED=0 \
	-e HOME=/home/gopher \
	-w /usr/src \
	$(BUILDIMAGE)

ifndef CIRCLECI
	DOCKERENV := docker $(DOCKERARGS)
else
	DOCKERENV :=
endif

.PHONY: all
all: generate build check ts docs

.PHONY: pull-build-image
pull-build-image: 
ifndef CIRCLECI
ifndef OFFLINE
	@docker pull $(BUILDIMAGE)
endif
endif

.PHONY: $(CACHEVOL)
$(CACHEVOL): pull-build-image Makefile
ifndef CIRCLECI
	@docker volume create $(CACHEVOL)
	@docker run -it --rm -v $(CACHEVOL):/usr/gocache \
		$(BUILDIMAGE) \
		sudo chown -R $(shell id -u):$(shell id -g) /usr/gocache
endif

.PHONY: $(MODVOL)
$(MODVOL): pull-build-image Makefile
ifndef CIRCLECI
	@docker volume create $(MODVOL)
	@docker run -it --rm -v $(MODVOL):/go/pkg/mod \
		$(BUILDIMAGE) \
		sudo chown -R $(shell id -u):$(shell id -g) /go/pkg/mod
endif

.PHONY: $(HOMEVOL)
$(HOMEVOL): pull-build-image Makefile
ifndef CIRCLECI
	@docker volume create $(HOMEVOL)
	@docker run -it 	--rm -v $(HOMEVOL):/home/gopher \
		-e GITHUB_TOKEN=$(GITHUB_TOKEN) \
		-e HOME=/home/gopher \
		$(BUILDIMAGE) \
		sudo chown -R $(shell id -u):$(shell id -g) /home/gopher
	@docker run -it --rm -v $(HOMEVOL):/home/gopher \
		-u $(shell id -u):$(shell id -g) \
		-e GITHUB_TOKEN=$(GITHUB_TOKEN) \
		-e HOME=/home/gopher \
		$(BUILDIMAGE) \
		configure-git
endif

# Generate go code for proto files
.PHONY: generate
generate: $(CACHEVOL) $(MODVOL) $(HOMEVOL)
	$(DOCKERENV) \
		go generate ./...

# Build go code 
.PHONY: build
build: generate
	cat go.mod
	go build ./...

# Check go code 
.PHONY: check
check: 
	zutano go check ./...

# Generate API docs
.PHONY: docs
docs: $(CACHEVOL) $(MODVOL) $(HOMEVOL)
	$(DOCKERENV) \
		protoc -I.:vendor:vendor/googleapis/:vendor/github.com/gogo/protobuf/protobuf/ \
			--doc_out=docs $(PROTOSOURCES) \
			--doc_opt=html,index-raw.html
	cat docs/header.txt docs/index-raw.html > docs/index.html
	rm docs/index-raw.html

# Generate API as typescript
.PHONY: ts
ts: $(CACHEVOL) $(MODVOL) $(HOMEVOL)
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
	go get github.com/stretchr/testify

check-version:
	zutano check api branch
