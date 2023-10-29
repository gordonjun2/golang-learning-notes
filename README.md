# golang-learning-notes
Contains notes and examples from my Golang programming learning journey

# Useful Videos
- [Golang Package Structure Explained in 7 Minutes](https://www.youtube.com/watch?v=1MdX9Z9fWWw)

# Useful Websites
- [Go CLI Commands](https://pkg.go.dev/cmd/go)
- [Tips to Create a Proper Go Project Layout](https://www.developer.com/languages/go-project-layout/)
- [Golang project structuring — Ben Johnson way](https://medium.com/sellerapp/golang-project-structuring-ben-johnson-way-2a11035f94bc)
- [Go Style Best Practices](https://google.github.io/styleguide/go/best-practices.html)
- [Go Cheat Sheet](https://github.com/a8m/golang-cheat-sheet)
- [Testing in Go: go test](https://ieftimov.com/posts/testing-in-go-go-test/)
- [Publishing a Go module](https://go.dev/doc/modules/publishing)
- [Checkers in Go](https://github.com/batkinson/checkers-go)
- [Excel2JSON in Go](https://github.com/gordonjun2/excel2json-go)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

# Steps to create a Golang project
- Create project directory with the proper project layout (refer to [here](https://www.developer.com/languages/go-project-layout/) or [here](https://medium.com/sellerapp/golang-project-structuring-ben-johnson-way-2a11035f94bc))
- Initialise new module in root of the project directory using
    ```
    go mod init [module-path]

    eg.
    go mod init github.com/gordonjun2/excel2json-go
    ```
- Write codes for the project
- Add dependencies to current project module and install them using
    ```
    go get [packages]

    eg.
    go get github.com/360EntSecGroup-Skylar/excelize
    ```
- Add missing and remove unused modules for the project module using
    ```
    go mod tidy
    ```
- To build the project module, build its executable using
    ```
    go build [packages]

    eg.
    go build ./cmd/xlsxlocal
    ```
    NOTE: If the path from the root of the directory to the entry .go file is *./cmd/xlsxlocal/main.go*, using ```go build ./cmd/xlsxlocal/main.go``` will cause a *main* executable to be generated whereas using ```go build ./cmd/xlsxlocal``` will cause a *xlsxlocal* executable to be generated. Both differ on the name of the executable generated.
- To run the generated module, use
    ```
    ./[executable-generated]

    eg.
    ./xlsxlocal
    ```
- To test the project module, use
    ```
    go test
    ```
    in the package folder
- To publish the project module, refer to [here](https://go.dev/doc/modules/publishing).

# Tips
- When using ```go test``` and the *go: cannot find main module* message is shown, this means that the current directory is not recognized as a Go module. To create one, run
```
go mod init <module-name>
```

