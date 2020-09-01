build:
	rm -rfv functions
	mkdir -p functions
	go get ./...
	go build -o functions ./...

	cp -r data functions/
