.PHONY: package

NAME:=rapid
MAINTAINER:="Sjoerd Siebinga <sjoerd@delving.eu>"
DESCRIPTION:="RAPID Linked Open Data Platform"
MODULE:=github.com/delving/rapid

GO ?= go
TEMPDIR:=$(shell mktemp -d)
VERSION:=$(shell sh -c 'grep "Version = \"" cmd/root.go  | cut -d\" -f2')
GOVERSION:=$(shell sh -c 'go version | cut -d " " -f3')

LDFLAGS:=-X main.Version=$(VERSION) -X main.BuildStamp=`date '+%Y-%m-%d_%I:%M:%S%p'` -X main.GitHash=`git rev-parse HEAD` -X main.BuildAgent=`git config user.email`

# var print rule
print-%  : ; @echo $* = $($*)

clean:
	rm -rf $(NAME) build report gin-bin result.bin *.coverprofile */*.coverprofile hub3/rapid.db hub3/models/rapid.db dist

clean-build:
	@make clean
	mkdir -p build

run:
	@go run main.go

build:
	@make clean-build
	@go build -a -o build/$(NAME) -ldflags=$(LDFLAGS) $(MODULE)

run-dev:
	gin -buildArgs "-i -ldflags '${LDFLAGS}'" run http

test:
	@go test  ./...

benchmark:
	@go test --bench=. -benchmem ./...

ginkgo:
	@ginkgo -r  -skipPackage go_tests

twatch:
	@ginkgo watch -r -skipPackage go_tests


compose-up:
	@docker-compose up

compose-down:
	@docker-compose down

compose-clean:
	@docker-compose down --volumes

release:
	@goreleaser --rm-dist --skip-publish
	@rpm --addsign dist/*.rpm
	@debsigs --sign=origin -k E2D6BD239452B1ED15CB99A66C417F6E7521731E dist/*.deb

release-dirty:
	@goreleaser --rm-dist --skip-publish --snapshot --skip-validate
	@rpm --addsign dist/*.rpm

release-snapshot:
	@goreleaser --rm-dist --skip-publish --snapshot
	@rpm --addsign dist/*.rpm

release-public:
	@goreleaser --rm-dist --skip-publish

dev-bootstrap:
	@go get -u github.com/codegangsta/gin
	@curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	@go get -u github.com/mattn/goveralls
	@go get -u github.com/goreleaser/goreleaser
	#@go get -u github.com/onsi/ginkgo/ginkgo
	#@go get github.com/onsi/gomega/..

goreport:
	@mkdir -p report
	@rm -rf report/*
	@goreporter -p ../rapid -r report -e vendor -f html
