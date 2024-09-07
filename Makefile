buf:
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf generate
	sudo chown -R 1000:1000 proto

build: buf
	go build -o sample_plugin main.go
	GOOS=wasip1 GOARCH=wasm go build -o sample_plugin.wasm main.go

.PHONY: buf sqlc
sqlc: build
	docker run --rm -v $(shell pwd):/src -w /src sqlc/sqlc:1.27.0 generate
	sudo chown -R 1000:1000 tutorial
