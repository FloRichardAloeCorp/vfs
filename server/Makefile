build_and_push_image:
	docker build . -t ghcr.io/florichardaloecorp/vfsserver:latest;
	docker push ghcr.io/florichardaloecorp/vfsserver:latest

format:
	go fmt ./...

upd-vendor:
	go mod tidy
	go mod vendor
 
test: upd-vendor 
	go clean -testcache
	GIN_MODE=release go test -timeout 1m -cover $$(go list ./... | grep -v test)

lint: 
	golangci-lint run --allow-parallel-runners -c ./.golangci-lint.yml --fix ./...