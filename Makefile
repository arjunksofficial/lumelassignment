swagger:
	swag init -g ./cmd/apiserver/main.go --output ./docs/swagger --pd --parseInternal --parseDepth 10  && rm ./docs/swagger/*.go