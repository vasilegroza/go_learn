# Functions [§](https://www.golang-book.com/books/intro/7) 

- A function is an independent section of code that maps zero or more input parameters to zero or more output parameters. 

    ![function](./function.png)
-  let's create a function average 
    ```go
    func average( elements []float64 ) float64 {
        panic ("Not implemented")
    }
    ```
- a few things to keep in mind
    - The names of the parameters don't have to match in the calling function. For example we could have done this:
    - The functions don't have access to anything in the calling function
    - Functions are build un in a **stack**

        ![stack](./stack.png)

##Returning Multiple Values [§](https://www.golang-book.com/books/intro/7#section2) 

- Go is capable of returning multiple values from a function

    ```go
    func f() (int, int){
        return 1, 2
    }
    ```
- Multiple values are often used to return an error value along with the result (x, err := f()), or a boolean to indicate success (x, ok := f()).

## Variadic Functions [§](https://www.golang-book.com/books/intro/7#section3)

- By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters.
    ```go
    func add (args ... int) int {
        sum := 0
        for _, value := range args {
            sum += value
        }
        return sum
    }
    ```

## Closure [§](https://www.golang-book.com/books/intro/7#section4)
- It is possible to create functions inside of functions:
    ```go
    func main() {
        add := func(x, y int) int {
            return x + y
        }
        fmt.Println(add(1,1))
    }
    ```
- When you create a local function like this it also has access to other local variables

- One way to use closure is by writing a function which returns another function which – when called – can generate a sequence of numbers. 


## Recursion [§](https://www.golang-book.com/books/intro/7#section5)
 - A function that is able to call itself.
       
    ```go
    func factorial(x uint) uint {
        if x == 0 {
            return 0
        }

        return x * factorial(x - 1)
    }
    ```
- Closure and recursion are powerful programming techniques which form the basis of a paradigm known as functional programming. 

## Deffer, Panic & Recover [§](https://www.golang-book.com/books/intro/7#section6)

- Go has a special statement called defer which schedules a function call to be run after the function completes.
    ```go
    package main

    import "fmt"

    func first() {
        fmt.Println("1st")
    }
    func second() {
        fmt.Println("2nd")
    }
    func main() {
        defer second()
        first()
    }
    ```
    is equal to 

    ```go
    func main() {
        first()
        second()
    }
    ```
- **defer** is often used when resources need to be freed in some way. For example when we **open** a file we need to make sure to **close** it later. With **defer**:
    ```go
    f, _ os.Opne(filename)
    defer f.Close()
    ```

- We can handle a run-time panic with the built-in **recover** function. recover stops the panic and returns the value that was passed to the call to panic

    ```go
    package main

    import "fmt"

    func main() {
    defer func() {
            str := recover()
            fmt.Println(str)
        }()
        panic("PANIC")
    }
    ```
## Problems
- [x] sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?
    - ```go
      func sum(numbers []int) int {
          panic("Not implemented")
      }  
      ```

- [x] Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).
    - here we go
    ```go
    func half(number int) bool {
        if (number % 2 ==1) {
            return false
        }
        return true
    }
    ```

- [x] Write a function with one variadic parameter that finds the greatest number in a list of numbers.
    - let's name it getMax(args ... int) int
    ```go
    func getMax(args ... int) (max int) {
        if len(args)==0 {
            panic("zero elements")
        }
        max = args[0]
        for _, value := range args {
            if value > max {
                max = value
            }
        }
        return max
    }
    ```

- [x] Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.
    - ```go
      func makeOddGenerator() func() uint {
          i := uint(1)
          return func () (rez uint){
              rez = i
              i += 2
              return rez
          }
      }
      ```
- [x] The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).
    - ```go
      func fib(n int) int {
        if(n == 0){
            return 0
        } else if (n == 1) {
            return 1
        }
        return fib(n - 1) + fib(n - 2)
          
      }
      ```

- [x] What are defer, panic and recover? How do you recover from a run-time panic? 
    - **defer** is a statement used to call a function at the end of current function
    - **panic & recover** with these two we can handle exceptions like try except in python