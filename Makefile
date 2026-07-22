.PHONY: gen push lint breaking

BUF ?= buf

gen:
	$(BUF) generate

push:
	$(BUF) push

lint:
	$(BUF) lint

breaking:
	$(BUF) breaking --against '.git#branch=main'