GENGO = go run ./internal/cmd/tool gen

gen.go:
	$(GENGO) ./pkg/apis/meta/v1

test:
	go test ./...

test.race:
	go test -race ./...

dep.update:
	go get -u ./...

fmt:
	gofumpt -w -l .
