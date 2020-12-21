.PHONY: setup
setup:
	GO111MODULE=off go get \
		golang.org/x/lint/golint \
		github.com/motemen/gobump/cmd/gobump \
		github.com/Songmu/make2help/cmd/make2help

.PHONY: lint
lint:
	go vet ./...
	golint -set_exit_status ./...
