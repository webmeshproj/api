LOCALBIN     ?= $(CURDIR)/bin
BUF          ?= $(LOCALBIN)/buf
BUF_VERSION  ?= 1.17.0

generate: gen

gen: buf update  ## Generate proto files.
	go install github.com/protobuf-tools/protoc-gen-deepcopy@latest
	rm -rf docs/
	$(BUF) generate proto
	npx typedoc --name "Webmesh API" --cleanOutputDir false

update: buf ## Update proto dependencies.
	cd proto/ ; $(BUF) mod update

.PHONY: buf
buf: $(BUF) ## Download buf locally if necessary.
$(BUF):
	mkdir -p $(LOCALBIN)
	test -s $(LOCALBIN)/buf || curl -sSL \
		https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(shell uname -s)-$(shell uname -m) \
		-o $(LOCALBIN)/buf && chmod +x $(LOCALBIN)/buf

publish: buf
	cd proto/ ; $(BUF) push
