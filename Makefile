ARCH=amd64
OS=linux
IMAGE=hub.global.cloud.sap/d067954/gotee
HASH := $(shell git rev-parse HEAD | head -c 7)
VERSION:=v$(shell date -u +%Y%m%d)-$(HASH)

build:
	GOOS=$(OS) GOARCH=$(ARCH) go build
	docker build -t $(IMAGE):$(VERSION) . 

push:
	docker push $(IMAGE):$(VERSION)
