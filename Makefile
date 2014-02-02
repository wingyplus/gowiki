deps:
	@go get -v ./...
deps-test:
	@go get -v github.com/onsi/ginkgo
	@go get -v github.com/onsi/gomega
test: deps deps-test
	@go test -v ./...
