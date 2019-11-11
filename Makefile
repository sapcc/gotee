GOFILES := $(wildcard *.go)

IMAGE=hub.global.cloud.sap/d067954/gotee
HASH := $(shell git rev-parse HEAD | head -c 7)
VERSION:=v$(shell date -u +%Y%m%d%M%S)-$(HASH)

#all: bin/gotee bin/gotee-linux

bin/gotee: $(GOFILES)
	go build -o bin/gotee

bin/gotee-linux: $(GOFILES)
	GOOS=linux GOARCH=amd64 go build -o bin/gotee-linux

docker: bin/gotee-linux
	docker build -t $(IMAGE):$(VERSION) . 
	docker push $(IMAGE):$(VERSION)

.Phony: clean
clean:
	rm -f bin/gotee bin/gotee-linux
