NETLIFY_FUNCTION_DIR = api

build:
	rm -rfv $(NETLIFY_FUNCTION_DIR)
	mkdir -p $(NETLIFY_FUNCTION_DIR)
	go get ./...
	go build -o $(NETLIFY_FUNCTION_DIR) ./...

	cp -r data $(NETLIFY_FUNCTION_DIR)/
