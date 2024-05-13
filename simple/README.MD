# go tools

## gotest
```go
go test -race -cover  -coverprofile=./coverage.out -timeout=10m -short -v ./...
go tool cover -func ./coverage.out
```

## godoc
```go
go get golang.org/x/tools/cmd/godoc
```