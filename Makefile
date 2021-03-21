.PHONY: assetutil test

assetutil:
	go build ./cmd/assetutil

test:
	go test ./... -cover
