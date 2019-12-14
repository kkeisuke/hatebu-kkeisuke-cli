COMMAND := htb
PACKAGE_NAME := github.com/kkeisuke/hatebu-kkeisuke-cli/api/algolia
APPLICATION := $(PACKAGE_NAME).application=$(ALGOLIA_APPLICATION)
API_KEY := $(PACKAGE_NAME).apiKey=$(ALGOLIA_API_KEY)
INDEX := $(PACKAGE_NAME).indexName=$(ALGOLIA_INDEX)

.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(COMMAND) -ldflags "-X $(APPLICATION) -X $(API_KEY) -X $(INDEX)"
