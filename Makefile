IMAGE_NAME=quotes-server
VERSION=0.0.1

build:
	go build -o ./bin/$(IMAGE_NAME) ./cmd/

run: build
	./bin/$(IMAGE_NAME)

image:
	docker build -t $(REGISTRY)/${IMAGE_NAME}:$(VERSION) -f ./Dockerfile .

dist:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/$(IMAGE_NAME) ./cmd/
	
scp: dist
	source .env && scp ./bin/$(IMAGE_NAME) ${REMOTE_SSH}: