ROOT = $(shell git rev-parse --show-toplevel)
CC_DIR := cc
CRD_Y := $(CC_DIR)/crd.y
CRD_GO := $(CC_DIR)/crd.go
CRD_OUTPUT := $(CC_DIR)/crd.output

.PHONY: build
build: generate dist/crd

.PHONY: generate
generate: go-regenerate $(CRD_GO)

$(CRD_GO): $(CRD_Y)
	goyacc -o $@ -v $(CRD_OUTPUT) $<

.PHONY: go-regenerate
go-regenerate: clean-go-generate go-generate

.PHONY: clean-go-generate
clean-go-generate:
	find $(ROOT) -name "*_generated.go" -type f | xargs rm -f

.PHONY: go-generate
go-generate:
	go generate ./...

.PHONY: prepare
prepare:
	go install github.com/berquerant/marker@latest
	go install golang.org/x/tools/cmd/stringer@latest

.PHONY: dist/crd
dist/crd:
	go build -o $@ cmd/crd/main.go

.PHONY: test
test:
	go test ./...
