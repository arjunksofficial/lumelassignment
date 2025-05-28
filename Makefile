swagger:
	swag init -g ./cmd/apiserver/main.go --output ./docs/swagger --pd --parseInternal --parseDepth 10  && rm ./docs/swagger/*.go


installmockery:
	go install github.com/vektra/mockery/v3@v3.3.1

generatemocks: installmockery
	rm -rf mocks && \
	export GOFLAGS=-buildvcs=false && \
	mockery 

coverage:
	go test -buildvcs=false -coverprofile=coverage.out ./...
	go tool cover -func coverage.out