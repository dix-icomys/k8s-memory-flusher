GO=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go
TAG=latest
BIN=flusher
BIN_PATH=build/$(BIN)
IMAGE=eugenelukin/k8s-memory-flusher

all: image
	docker push $(IMAGE):$(TAG)

build: clean
	$(GO) build -o $(BIN_PATH) .

image: build
	docker build -t $(IMAGE):$(TAG) .

clean:
	rm -f $(BIN_PATH)
