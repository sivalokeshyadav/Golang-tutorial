# Slides for daily sessions

## Day 1

### Agenda

- Intro
- Topics to be covered
- Why Go?
- History
- Go playground
- Hello world example
- Hello web example

### Environment setup

- Install Go
- Install Visual Studio Code and Analysis Tools

## Day 2

### Git basics

- git commit is for your local repository
- git push is to origin (remote repository)
- git clone is to clone the repo to local

### To do

- Create a github account
- install git in your system
- Create SSH pub / private keys

```ssh-keygen -t ed25519 -C "your_email@example.com"```

- add the pub key to github.com at settings
- cheat sheets for reference

https://training.github.com/downloads/github-git-cheat-sheet
https://about.gitlab.com/images/press/git-cheat-sheet.pdf

### Understading the basic program

- The three main areas to fous
  - package main-> every file needs to define a package, package main is mandatory for any program to run
  - import -> to get and use other packages, only needed if you are using other packages
  - function main()-> you need a main() function to run the program

### Understanding packages

- main package -> Executable package, every program needs a main package if we have to run the program. Only the main packages compiles and generates binaries (ex: main.exe)
- reusable packages -> these the packages to share / reuse the code between modules
- Typically, the folder name would be the package name

### Go CLI

- Command line interface, you get the cli when you install go.
- You use this cli to work with Go
  - `go build` - This will compiles and generates the binary. You need a main package to generate binary
  - `go run` - This will build and run the binary
  - `go fmt` - Formats the code
  - `go install` - To install a go library / tool
  - `go get` - To get a package to our system
  - `go test` - To run the unit tests

## Day 3

### Buit in types

- bool, string
- int int8 int16 int32 int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // alias for uint8
- rune // alias for int32
- float32 float64
- complex64 complex128

### Variables

- `var` is the keyword
- Varibles can be decaled once
- make sure to use `:=` to declare and initalize
- Go is a strictly typed lang -> you can't change the type once declared

### Zero values

- Variables declared without an explicit initial value are given their zero value.
- `0` for numeric types,
- `false` for the boolean type
- `""` (the empty string) for strings

| Type          | Zero Value  |
| -----------   | ----------- |
| Integer       | 0           |
| Floating point| 0.0         |
| Boolean       | false       |
| String        | ""          |
| Pointer       | nil         |
| Interface     | nil         |
| Slice         | nil         |
| Map           | nil         |
| Channel       | nil         |
| Function      | nil         |

### Constants

- used with `const` keyword
- It's even if a const is declared and not used

## Day 4

### Type conversions

- For converting compatible types you just use `new_type(old_value)`

### Functions

- Go supports first class functions
- different ways of creating and using the functions in examples

### Use Case Playing cards

- Deck of cards functionality
  - New Deck
  - Print Deck
  - Shuffle
  - Deal
  - Save to file
  - Read from file

### Assignement 1

- Return a new deck of cards with suits and numbers

## Day 5

### Slices / Arrays

- Array is a data structure to maintain a list of similar types
- Length is fixed for arrays in Go
- Index starts from 0
- how to create an array?
- Slices are reference data types which points to underlying array
- Array is a value data type
- Slice is Reference data type

### Flow control statements

- if / else
- for
- switch
- defer
- blank reference?

### Assignement 2

- Complete the shuffle function for deck of cards

## Day 6

### Use Case cont1

- deal func
- save to file
- assignment for read from file

### Structs

- What is Struct?
- Definining struct
- Declaration / Initialization
- Embdedded structs
- Reciver functions / methods

### Functions / Methods

### Assignement 3

- Read the deck of cards from file

## Day 7

### Use Case cont2

- Shuffle
- Read from file

### Pointers

- Use `&` to access the pointer (address)
- Use `*` to access the value from pointer
- `*` in front for a type to indicate it's a pointer type
- Pointer indirection

### Value Types / Reference Types

| Value Type       | Referece Type |
| ---------------- | ------------- |
| int              | slices        |
| float            | maps          |
| string           | channels      |
| bool             | pointers      |
| structs          | functions     |
| arrays           |      |

### Basics of go testing

- Unit testing -> to test a piece of code be it a method or a function

### Assignement 4

- Write some unit tests when writing a file and reading file

## Day 8

### Empty Struts

- Empty structs take 0 memory
- Empty structs always points to the same location (address)
- Useful to define methods for a service

### Mpas

- All the keys should be of same type
- All the values should be of same type
- Map Vs Structs

| Map                 | Struct               |
|-------------------- |----------------------|
| All keys - same type| different types      |
| related properties  | represent something    |
| All values - same type|can be different |
| keys - can be added | Define them at compile time |
| keys are indexed   | keys not indexed |
| Reference Type     | Value Type |

### make vs new

- The new built-in function allocates memory, the value returned is a pointer to a newly allocated zero value of that type.
- The make built-in function allocates and initializes an object of type slice, map, or chan (only).

### Assignment 5

- Print numbers from 1 to 10 and print whether it's even or odd
- Sample output

``` text
1 - odd
2 - even
3 - odd
4 - even
...
```

## Day 9

- Map beceomes a Set when value is an empty Struct

### Interfaces

- What is the interface in OOP?
An interface defines the behavior of an object. It only specifies what the object is supposed to do. It is actually a concept of abstraction and encapsulation
- To implment an interface
  - No need to explicitly mention that it implements
  - No `impliments` or `extends` keywords need to say the struct implment any interface
- If you define an interface in the file... All the structs in the file automatically be the members of that interface
  - So you must honor and impliment all the methods in that interface to use that struct as interface
- Check the example for interfaces

### Assignment 6

- Add a incomestream type `RentalIncome` to `.\examples\interfaces\main.go`
  - Rental income will have
    - Property Name
    - Monthly rent
    - No of months

## Day 10

### Channels

- Channels are a typed conduit through which you can send and receive values with the channel operator `<-`
- `ch <- v`    // Send v to channel ch
- `v := <-ch`  // Receive from ch, and assign value to v
- Channels are created for a single type
- non-buffered channels
  - 0 capcity
  - used for synchronous communication
  - You should have some goRoutine ready to recieve the data
- Buffered channels
  - capacity is defined
  - error when you try to push to channel which is full
  - async communication

### Select / Switch

- Switch -> only executes the first matching case
- Select
  - It is only used for Channels
  - It checks for multiple macthing cases
  - If more than one is true then it will choose random
  - It can be used to write to channel as well

### Assignment 7

- Use `select` to print the fibnocci series
- Should use a secondary channel as quit signal

## Day 11

### GOPATH / GOROOT

- GOROOT is a variable that defines where your Go SDK is located. You do not need to change this variable, unless you plan to use different Go versions.
- GOPATH is a variable that defines the root of your workspace. By default, the workspace directory is a directory that is named go within your user home directory (~/go for Linux and MacOS, %USERPROFILE%/go for Windows). GOPATH stores your code base and all the files that are necessary for your development. You can use another directory as your workspace by configuring GOPATH for different scopes. GOPATH is the root of your workspace and contains the following folders:
  - src/: location of Go source code (for example, .go, .c, .g, .s).
  - pkg/: location of compiled package code (for example, .a).
  - bin/: location of compiled executable programs built by Go.

### GoLang workspace

- Define you GOPATH
- Be sure of your GOROOT
- git@github.com:rajesh4b8/users-api-batch4.git

### MVC Architecture

- MVC is abbreviated as Model View Controller is a design pattern created for developing applications specifically web applications. As the name suggests, it has three major parts. The traditional software design pattern works in an "Input - Process - Output" pattern whereas MVC works as "Controller -Model - View" approach

### Use Case - Users API

- Create User
- Read / List Users
- Update User
- Delete User

## Day - 13

### RESTFul web services

- To communicate over web with http protocol
- Different http methods
  - GET
  - POST
  - PUT
  - PATCH
  - DELETE

### Modules

- GOROOT -> Where Go is installed (default for windows: `C:\Program Files\Go`)
- GOPATH -> Where your go workspace is defined (Default for windows: `C:\Users\<user>\Go`)
- All the modules being used will be downloded to GOPATH
- Go Modules introduced in Go 1.13
- If you define go modules to the package you nolonger need to put your repository in GOPATH
- Modules are how Go manages dependencies.
- A module is a collection of packages that are released, versioned, and distributed together. Modules may be downloaded directly from version control repositories or from module proxy servers.

## Day - 19

### Go-kit

- https://gokit.io/faq/#what-is-go-kit
- example