THIS_FILE := $(lastword $(MAKEFILE_LIST))
NETLIFY_FUNCTION_DIR = functions

.PHONY: build
build:
	@$(MAKE) -f $(THIS_FILE) clean
	-mkdir -p $(NETLIFY_FUNCTION_DIR)
	
	go generate src/films/utils/films-handler.go
	go generate src/people/utils/people-handler.go
	go get ./...
	go build -o $(NETLIFY_FUNCTION_DIR) ./...


.PHONY: clean
clean:
	-rm -rfv $(NETLIFY_FUNCTION_DIR)
	-rm src/films/utils/generated-films.go
	-rm src/people/utils/generated-people.go