THIS_FILE := $(lastword $(MAKEFILE_LIST))
NETLIFY_FUNCTION_DIR = functions

.PHONY: build
build:
	@$(MAKE) -f $(THIS_FILE) clean
	-mkdir -p $(NETLIFY_FUNCTION_DIR)
	
	go generate src/api.go
	go get ./...
	go build -o $(NETLIFY_FUNCTION_DIR) ./...
	# renaming the executable so that on netlify we can access it with /api/
	mv $(NETLIFY_FUNCTION_DIR)/src $(NETLIFY_FUNCTION_DIR)/api


.PHONY: clean
clean:
	-rm -rfv $(NETLIFY_FUNCTION_DIR)
	-rm -rfv generated
