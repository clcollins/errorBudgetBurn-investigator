# List all available make targets 
.PHONY: list 
list: 
	@LC_ALL=C $(MAKE) -pRrq -f $(firstword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/(^|\n)# Files(\n|$$)/,/(^|\n)# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | grep -E -v -e '^[^[:alnum:]]' -e '^$@$$' 
 
.PHONY: install-tools
install-tools: install-golangci-lint install-goreleaser install-cobra-cli install-addlicense

.PHONY: install-golangci-lint
install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: install-goreleaser
install-goreleaser:
	go install github.com/goreleaser/goreleaser@latest

.PHONY: install-cobra-cli
install-cobra-cli:
	go install github.com/spf13/cobra-cli@latest

.PHONY: install-addlicense
install-addlicense:
	go install github.com/google/addlicense@latest