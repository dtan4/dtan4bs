NAME     := dtan4bs
VERSION  := v0.1.0
REVISION := $(shell git rev-parse --short HEAD)

SRCS     := $(shell find . -type f -name '*.go')
LDFLAGS  := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
NOVENDOR := $(shell go list ./... | grep -v vendor)

DIST_DIRS := find * -type d -exec

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): $(SRCS)
	go generate $(NOVENDOR)
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: ci-test
ci-test:
	echo "" > coverage.txt
	set -e; \
	for d in $(NOVENDOR); do \
		go test -coverprofile=profile.out -covermode=atomic -v $$d; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi; \
	done

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: cross-build
cross-build: deps
	set -e; \
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o dist/$$os-$$arch/$(NAME); \
		done; \
	done

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u -v github.com/golang/dep/cmd/dep
endif

.PHONY: deps
deps: dep go-bindata
	dep ensure -v

.PHONY: dist
dist:
	cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar -zcf $(NAME)-$(VERSION)-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r $(NAME)-$(VERSION)-{}.zip {} \; && \
	cd ..

.PHONY: go-bindata
ifeq ($(shell command -v go-bindata 2> /dev/null),)
	go get -u -v github.com/jteeuwen/go-bindata/...
endif

.PHONY: install
install:
	go generate $(NOVENDOR)
	go install $(LDFLAGS)

.PHONY: release
release:
	git tag $(VERSION)
	git push origin $(VERSION)

.PHONY: test
test:
	go generate $(NOVENDOR)
	go test -cover -v $(NOVENDOR)

.PHONY: update-deps
update-deps: dep
	dep ensure -update -v
