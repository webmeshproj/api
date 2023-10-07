LOCALBIN     ?= $(CURDIR)/bin
BUF          ?= $(LOCALBIN)/buf
BUF_VERSION  ?= 1.17.0

generate: gen

gen: buf ## Generate proto files.
	go install github.com/protobuf-tools/protoc-gen-deepcopy@latest
	$(BUF) generate proto
	pandoc -f html -t gfm -o README.md doc/index.html

.PHONY: buf
buf: $(BUF) ## Download buf locally if necessary.
$(BUF):
	mkdir -p $(LOCALBIN)
	test -s $(LOCALBIN)/buf || curl -sSL \
		https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(shell uname -s)-$(shell uname -m) \
		-o $(LOCALBIN)/buf && chmod +x $(LOCALBIN)/buf

publish: buf
	cd proto/ ; $(BUF) push