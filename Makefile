# Copyright 2017 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

TARG.Name:=openpitrix
TRAG.Gopkg:=openpitrix.io/openpitrix

DOCKER_TAGS=latest

GO:=docker run --rm -it -v `pwd`:/go/src/$(TRAG.Gopkg) -w /go/src/$(TRAG.Gopkg) golang:1.9-alpine go

MYSQL_DATABASE:=$(TARG.Name)
MYSQL_ROOT_PASSWORD:=password

.PHONY: help
help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  all               to generate, test and release"
	@echo "  start             to start service (port:9100)"
	@echo "  stop              to stop service"
	@echo "  tools             to install depends tools"
	@echo "  init-vendor       to init vendor packages"
	@echo "  update-vendor     to update vendor packages"
	@echo "  generate          to generate restapi code"
	@echo "  build             to build the service"
	@echo "  test              to run go test ./..."
	@echo "  release           to build and release current version"
	@echo "  clean             to clean the temp files"

.PHONY: all
all: generate test release

.PHONY: init-vendor
init-vendor:
	@if [[ ! -f "$$(which govendor)" ]]; then \
		go get -u github.com/kardianos/govendor; \
	fi
	govendor init
	govendor add +external
	@echo "ok"

.PHONY: update-vendor
update-vendor:
	@if [[ ! -f "$$(which govendor)" ]]; then \
		go get -u github.com/kardianos/govendor; \
	fi
	govendor update +external
	govendor list
	@echo "ok"

.PHONY: tools
tools:
	docker pull golang:1.9-alpine
	docker pull openpitrix/openpitrix:api-builder
	docker pull mysql
	@echo "ok"

.PHONY: generate
generate:
	cd ./api && make generate
	cd ./pkg/cmd/api && make
	go generate ./pkg/version/

.PHONY: mysql-start
mysql-start:
	@docker run --rm --name openpitrix-db -e MYSQL_ROOT_PASSWORD=$(MYSQL_ROOT_PASSWORD) -e MYSQL_DATABASE=$(MYSQL_DATABASE) -p 3306:3306 -d mysql:5.6 || docker start openpitrix-db
	@echo "ok"

.PHONY: mysql-stop
mysql-stop:
	@docker stop openpitrix-db
	@echo "ok"


.PHONY: build
build: generate
	$(GO) fmt ./...
	docker build -t $(TARG.Name) -f ./Dockerfile .
	@docker image prune -f 1>/dev/null 2>&1
	@echo "ok"

.PHONY: start
start:
	docker-compose up -d
	@echo "ok"

.PHONY: stop
stop:
	docker-compose down
	@echo "ok"

.PHONY: release
release:
	@echo "TODO"

.PHONY: test
test:
	$(GO) fmt ./...
	$(GO) vet ./...
	$(GO) test ./...

	docker-compose up --build -d
	sleep 10

	curl localhost:9100/ping
	curl localhost:9100/panic

	curl localhost:9100/v1/apps
	curl localhost:9100/v1/appruntimes
	curl localhost:9100/v1/clusters
	curl localhost:9100/v1/repos

	curl localhost:9100/v1/apps/invalid-id
	curl localhost:9100/v1/appruntimes/invalid-id
	curl localhost:9100/v1/clusters/invalid-id
	curl localhost:9100/v1/repos/invalid-id

	curl localhost:9100/v1/apps/app-panic000
	curl localhost:9100/v1/appruntimes/rt-panic000
	curl localhost:9100/v1/clusters/cl-panic000
	curl localhost:9100/v1/repos/repo-panic000

	docker-compose down
	@echo "ok"

.PHONY: clean
clean:
	@echo "ok"
