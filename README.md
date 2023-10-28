# golang
Learn golang


# Go CLI

- Compiles a bunch of go source code files
    ```
    $ go build
    ```
- Compiles and executes one or two files
    ```
    $ go run
    ```

- Formats all the code in each file in the current directory
    ```
    $ go fmt
    ```
- Compiles and "installs" a package
    ```
    $ go install
    ```

- Downloads the raw source code of someone else's package
    ```
    $ go get
    ```
- Runs any tests associated with the current project
    ```
    $ go test
    ```
# Package == Project == Workspace
- https://golang.org/pkg/
- package name must be declared in every file of the package
    ```
    package main
    ```
- 2 types of packages
    - Executable
        - Generate a file that we can run
        - "main" is used to create executable package that generates main.exe
    - Reusable
        - Code used as 'helpers'
        - Other than "main" creates reusable package
    

# import "library name or resusable package name"
- Imports library or reusable package
- "fmt" is part of standard library



