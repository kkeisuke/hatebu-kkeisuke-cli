COMMAND := htb
PACKAGE_NAME := github.com/kkeisuke/hatebu-kkeisuke-cli/api/algolia
APPLICATION := $(PACKAGE_NAME).application=$(ALGOLIA_APPLICATION)
API_KEY := $(PACKAGE_NAME).apiKey=$(ALGOLIA_API_KEY)
INDEX := $(PACKAGE_NAME).indexName=$(ALGOLIA_INDEX)

.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(COMMAND) -ldflags "-X $(APPLICATION) -X $(API_KEY) -X $(INDEX)"

.PHONY: test
test:
	go test -v -cover -coverpkg=github.com/kkeisuke/hatebu-kkeisuke-cli/api/algolia github.com/kkeisuke/hatebu-kkeisuke-cli/api/algolia_test
	go test -v -cover -coverpkg=github.com/kkeisuke/hatebu-kkeisuke-cli/domain/service github.com/kkeisuke/hatebu-kkeisuke-cli/domain/service_test
	go test -v -cover -coverpkg=github.com/kkeisuke/hatebu-kkeisuke-cli/domain/value github.com/kkeisuke/hatebu-kkeisuke-cli/domain/value_test
	go test -v -cover

.PHONY: lint
lint:
	go vet ./...; golint -set_exit_status ./...

.PHONY: setup-ci
setup-ci:
	GOBIN=$(pwd)/bin go get golang.org/x/lint/golint

.PHONY: lint-ci
lint-ci:
	go vet ./...; bin/golint -set_exit_status ./...
