u := $(if $(update),-u)

.PHONY: deps
deps:
	go get ${u} -d
	go mod tidy

.PHONY: devel-deps
devel-deps: deps
	go install github.com/Songmu/godzil/cmd/godzil@latest

.PHONY: test
test: deps
	go test

.PHONY: devel-deps
release: devel-deps
	godzil release
