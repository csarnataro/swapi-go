NETLIFY_FUNCTION_DIR = functions

build:
	rm -rfv $(NETLIFY_FUNCTION_DIR)
	mkdir -p $(NETLIFY_FUNCTION_DIR)
	go get ./...
	go build -o $(NETLIFY_FUNCTION_DIR) ./...

	echo "data folder (original)"
	ls -l $(NETLIFY_FUNCTION_DIR)/data
	cp -r data $(NETLIFY_FUNCTION_DIR)/
	echo "root folder"
	ls -l $(NETLIFY_FUNCTION_DIR)
	echo "data folder"
	ls -l $(NETLIFY_FUNCTION_DIR)/data
