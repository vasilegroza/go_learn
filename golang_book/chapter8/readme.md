# Pointers [§](https://www.golang-book.com/books/intro/8) 

- Pointers reference a location in memory where a value is stored rather than the value itself. 

## The **\*** and **&** operators [§](https://www.golang-book.com/books/intro/8#section1) 
- In Go a pointer is represented using the * (asterisk) character followed by the type of the stored value.
    ```go
    var p *int
    *p = 0
    ```
- When we write *p = 0 we are saying “store the int 0 in the memory location p refers to”. If we try p = 0 instead we will get a compiler error

- **&** returns the address of a variable.
    ```go
    var x int = 10
    p := &x

    ```
 - &x returns a *int (pointer to an int) because x is an int. 

## new [§]()
 - Another way to get a pointer is to use the built-in new function:
    ```go
    p := new(int)
    *p = 1
    fmt.Println(*p) //will print 1
    ```
- **new** takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.

## Problems
- [x] How do you get the memory address of a variable?
    - &variable_name (returns the address of a variable)

- [x] How do you assign a value to a pointer?
    -let's say we have p *int
    ```go
    var p *int
    *p = 5 //will assign value 5 at the address that p points to
    ```
- [x] How do you create a new pointer?
    ```go
    var p *int
    p1 := new(int)
    ```
- [x] What is the value of x after running this program:
    ```go
    func square(x *float64) {
        *x = *x * *x
    }
    func main() {
        x := 1.5
        square(&x)
    }
    ```
    - 1.5*1.5 = 2.25

- [x] Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).
    - see [solution](./main.go)