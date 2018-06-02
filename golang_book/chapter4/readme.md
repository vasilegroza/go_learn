# Variables

- A variable is a storage location, with a specific type and an associated name. 

- examples:
    ```go
    var message string = "hi there"
    var name string
    name = "human"
    welcome := "Welcome to our world"

    fmt.Println(message, name, "!" ,'\n', welcome)
    ```
## How to Name a Variable

- Naming a variable properly is an important part of software development.
- Names must start with a letter and may contain letters, numbers or the _ (underscore) symbol
- Pick names which clearly describe the variable's purpose.

## Scope
 - local
    ```go
    package main

    import "fmt"

    func main() {
        var x string = "Hello World"
        fmt.Println(x)
    }
    func f() {
        //fmt.Println(x) this will give error as the x is not defined in the current block aka { }
    }
    ```
 - global - all can have access to the variable
   ```go
   var x string = "Hello World"

    func main() {
        fmt.Println(x)
    }

    func f() {
        fmt.Println(x)
    }
   ```
## Constants

Go also has support for constants. Constants are basically variables whose values cannot be changed later. 

- instead of using var keyword use const
    ```go
    package main
    import fmt

    func main() {
        const hi_message = "Hello"
        fmt.Println(hi_message)
    }
    ```

## Defining Multiple Variables
Go also has another shorthand when you need to define multiple variables:

```go
var (
    a = 3
    b = 4
    c = 5
)
```

# Problems

* [x] What are two ways to create a new variable?
    - full definition   
        ```go
        var number int = 10
        ```
    - shorthand 
        ```go
        number := 10
        ```
* [x] What is the value of x after running: 
x := 5; x += 1?
    - the value of is would be **6**

* [x] What is scope and how do you determine the scope of a variable in Go?
    - The scope of a variable is the immediate block in which a variable was declared.

* [x] What is the difference between var and const?
    - the variable can be changed after it's initialization and the constant can't.

* [x] Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)
    - see [convertor](./convertor/main.go)
* [x] Write another program that converts from feet into meters. (1 ft = 0.3048 m)
    - see [convertor](./convertor.1/main.go)
    