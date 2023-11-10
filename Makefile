LOCALBIN  ?= $(CURDIR)/bin
BUF       ?= buf
VERSION   ?= $(shell git describe --tags --always)

generate: gen

gen: update  ## Generate proto files.
	go install github.com/protobuf-tools/protoc-gen-deepcopy@latest
	rm -rf docs/
	$(BUF) generate proto
	go run cmd/tsutil-gen/main.go --out ts/utils/rpcdb.ts
	cd ts && npx typedoc --name "Webmesh API" --cleanOutputDir false

update: ## Update proto dependencies.
	cd proto/ ; $(BUF) mod update

publish:
	cd proto/ ; $(BUF) push

npm-publish:
	sed -i 's/"version": ".*"/"version": "$(subst v,,$(VERSION))"/' ts/package.json
	cd ts && npm publish --access public