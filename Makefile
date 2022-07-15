go := $(shell command -v go 2>/dev/null)

controller-gen: version := v0.9.2
controller-gen:
	$(go) install sigs.k8s.io/controller-tools/cmd/controller-gen@$(version)

generate:
	$(go) generate ./...
