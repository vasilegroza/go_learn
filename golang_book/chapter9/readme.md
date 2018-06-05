# Structs and Interfaces [§](https://www.golang-book.com/books/intro/9)

## Structs  [§](https://www.golang-book.com/books/intro/9#section1)

- A struct is a type which contains named fields. For example we could represent a Circle like this:
    ```go
    type Circle struct {
        x float64
        y float64
        r float64
    }
    ```
- The type keyword introduces a new type. It's followed by the name of the type (Circle), the keyword struct to indicate that we are defining a struct type and a list of fields inside of curly braces. Each field has a name and a type.

### Initialization 
- se can a instance of Circle in many ways 
    ```go
    var c Circle // this will set struct at 0 for struct zero means 0 for each it's field
    c1 := new(Circle) //This allocates memory for all the fields, sets each of them to their zero value and returns a pointer. (*Circle)
    c2 := Circle{0, 0, 5}
    ```
## Fields
 - We can access field by using ' **.** ' operator
    ```go
    c := Circle{1, 1, 5}
    fmt.Println(c.x, c.y, c.r)
    c.x = 5
    c.y = 5
    ```
 - or compute area with a function
    ```go
    func circleArea(c *Circle) float64 {
        return math.Pi * c.r*c.r
    }
    ```
## Methods [§](https://www.golang-book.com/books/intro/9#section2)

 ```go
 func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
 }
 ```
 - In between the keyword func and the name of the function we've added a “receiver”. The receiver is like a parameter – it has a name and a type – but by creating the function in this way it allows us to call the function using the . operator
  ```go
  c := Circle{0, 0, 5}
  fmt.Println(c.area())
  ```
 - This is much easier to read, we no longer need the & operator (Go automatically knows to pass a pointer to the circle for this method) 

## Interfaces
 - The Circle and the Rectangle structures have in common the method area so let us create Shape interface

    ```go
    type Shape interface {
        area() float64
    }
    ```
 - Like a struct an interface is created using the type keyword, followed by a name and the keyword interface. But instead of defining fields, we define a “method set”. A method set is a list of methods that a type must have in order to “implement” the interface.
 s
 - then we can do this
 ```go
 func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
        area += s.area()
    }
    return area
 }
 ```

# Problems
- [x] What's the difference between a method and a function? 
    - a method is function that is belongs to a structure


- [x] Why would you use an embedded anonymous field instead of a normal named field?
    - when you want to reuse existing field and methods of other structures

- [x] Add a new method to the Shape interface called perimeter which calculates the perimeter of a shape. Implement the method for Circle and Rectangle.
    - see [solution](./main.go)