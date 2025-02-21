gen:
	go generate ./...

test:
	go test ./...

test.race:
	go test -race ./...

dep.update:
	go get -u ./...

fmt:
	go tool gofumpt -w -l .
