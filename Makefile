vendor:
	go mod vendor

test: vendor
	go test