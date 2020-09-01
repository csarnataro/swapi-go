NETLIFY_FUNCTION_DIR = functions

build:
	rm -rfv $(NETLIFY_FUNCTION_DIR)
	mkdir -p $(NETLIFY_FUNCTION_DIR)
	cp -r /opt/build/repo/data $(NETLIFY_FUNCTION_DIR)
	go get ./...
	go build -o $(NETLIFY_FUNCTION_DIR) ./...

	ls -lRa /opt/build/repo
	ls -lRa /var/task
