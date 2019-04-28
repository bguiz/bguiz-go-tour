# Learning Go

## General

- Setup:
  - https://golang.org/doc/install
  - `go version` - make sure it outputs `go1.12.4` or higher
- Modules
  - https://blog.golang.org/using-go-modules
- Concurrency
  - [Rob Pike - 'Concurrency Is Not Parallelism'](https://www.youtube.com/watch?v=cN_DpYBzKso)
- Testing
  - https://github.com/golang/go/wiki/TableDrivenTests
- Development tools
  - General
    - Localhost documentation:
      - `go get golang.org/x/tools/cmd/godoc`
      - `godoc -http :9111` and then visit http://localhost:9111/pkg/ in your broswer
    - IDE support: `go get -u -v github.com/mdempsky/gocode`
    - IDE debugging: `go get -v github.com/go-delve/delve/cmd/dlv`
    - Check for when you have not handled errors in your tests: `go get -u github.com/kisielk/errcheck`
  - For VS Code (IDE)
    - IDE extension
      - https://github.com/Microsoft/vscode-go
      - `code --install-extension ms-vscode.go`

## Learning

- Go Tour:
  - https://tour.golang.org/
  - Found in: `01-go-tour`
  - following along with the examples in the tour
  - For each file, enter `go run $FILENAME`,
    to interpret without producing a binary file on disk
- Go by Example
  - https://gobyexample.com/
- Learn go with Tests
  - https://github.com/quii/learn-go-with-tests
- Go Koans
  - https://github.com/cdarwin/go-koans

## Author

[Brendan Graetz](http://bguiz.com)

## Licence

GPL-3.0
