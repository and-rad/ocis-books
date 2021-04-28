# compile binaries & move them into output directory
CGO_ENABLED=0 go build -o bin/books github.com/and-rad/ocis-books/cmd/books
