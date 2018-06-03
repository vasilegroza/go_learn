# Control Structures [§](https://www.golang-book.com/books/intro/5)

## For [§](https://www.golang-book.com/books/intro/5#section1)
- The for statement allows us to repeat a list of statements (a block) multiple times. 
```go
for i:=0; i<10; i++ {
    fmt.Println(i)
}
```
- Other programming languages have a lot of different types of loops (while, do, until, foreach, …) but Go only has one that can be used in a variety of different ways. 

## If [§](https://www.golang-book.com/books/intro/5#section2)
- the basic format is:
    ```go
    if condition {

    } else {

    }
    ```
- Let's check for the previous elements whether they are even or not
    ```go
    if i%2 == 0 {
        //even
    } else {
        //odd
    }

    ```
- If statements can also have else if parts:
    ```go
    if i % 2 == 0 {
    // divisible by 2
    } else if i % 3 == 0 {
    // divisible by 3
    } else if i % 4 == 0 {
    // divisible by 4
    }
    ```
## Switch [§](https://www.golang-book.com/books/intro/5#section3)

- More elegant than elseif if there are more than 3 cases
    ```go
    switch i {
        case 0: fmt.Println("Zero")
        case 1: fmt.Println("One")
        case 2: fmt.Println("Two")
        case 3: fmt.Println("Three")
        case 4: fmt.Println("Four")
        case 5: fmt.Println("Five")
        default: fmt.Println("Unknown Number")
    }
    ```

# Problems
* [x] What does the following program print:
    
    ```go
    i := 10
    if i > 10 {
        fmt.Println("Big")
    } else {
        fmt.Println("Small")
    }
    ```
  - the above sequence prints **Small**

* [x] Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)
    - see [solution](./divisible_by_3/main.go)

* [x] Write a program that prints the numbers from 1 to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz". For numbers which are multiples of both three and five print "FizzBuzz".
    - see [solution](./fizbuz/main.go)