build:
	mkdir -p functions/api
	go get ./...
	go build -o functions/api ./...

