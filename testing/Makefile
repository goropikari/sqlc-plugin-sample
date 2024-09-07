build:
	go build -o sqlc-dump-proto main.go
	docker run --rm -v $(shell pwd):/src -w /src sqlc/sqlc:1.27.0 generate
	sudo chown -R 1000:1000 gen
