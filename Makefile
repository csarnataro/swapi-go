NETLIFY_FUNCTION_DIR = functions

build:
	rm -rfv $(NETLIFY_FUNCTION_DIR)
	mkdir -p $(NETLIFY_FUNCTION_DIR)
	cp -r /opt/build/repo/data $(NETLIFY_FUNCTION_DIR)
	go get ./...
	go build -o $(NETLIFY_FUNCTION_DIR) ./...

	mv -r /opt/build/repo/data mv -r /var/task/data
	echo "root folder"
	ls -lr /var/task
