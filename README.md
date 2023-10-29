# golang
Not object oriented



### Go CLI

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
### Package == Project == Workspace
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
    

### import "library name or resusable package name"
- Imports library or reusable package
- "fmt" is part of standard library



### Types & Variable
- Golang is strongly and statically typed.
- Type is required
- Basic types
    - bool
    - string
    - int
    - float64
- :=
    - card := "I'm a string"
    - Asks compiler to figure out what type of variable is card.
    - Only used when initializing a variable

### Extended types
- type deck[] string

### Function
- func func_name() return_type { ... }
    - func newCard() string

### Data structure
- Array
    - Fixed length list
- Slice
    - An array that can grow or shrink
    - All elements must be of same type
    - Initialize
        - cards := []string{newCard(), newCard()}
    - Append
        - cards = append(cards, "Six of Spades")

### Loop
- Enumerate
    - for i, card := range cards { ... }
- Index
    - for i := range d { ... }

### Object Oriented
- Receiver function
    - func (d deck) print() { ... }
    - Receivers sets up methods on variables that we create
- Normal functions in same file for better visualization

### Test
- go mod init package_name
- Test file should end with _test.go
    - deck_test.go
- To execute tests in a package
    - go test
