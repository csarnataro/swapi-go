NETLIFY_FUNCTION_DIR = functions

build:
	rm -rfv $(NETLIFY_FUNCTION_DIR)
	mkdir -p $(NETLIFY_FUNCTION_DIR)
	go get ./...
	go build -o $(NETLIFY_FUNCTION_DIR) ./...

	cp -r /opt/build/repo/functions/data $(NETLIFY_FUNCTION_DIR)/
	echo "root folder"
	ls -l $(NETLIFY_FUNCTION_DIR)
