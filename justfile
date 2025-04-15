gen:
    go generate ./...

test:
    go test ./...

test-race:
    go test -race ./...

update:
    go get -u ./...

dep:
    go mod tidy

fmt:
    go tool gofumpt -w -l .
