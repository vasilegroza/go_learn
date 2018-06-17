# Packages [§](https://www.golang-book.com/books/intro/11)

- Go was designed to be a language that encourages good software engineering practices. An important part of high quality software is code reuse – embodied in the principle “Don't Repeat Yourself.”


## Creating Packages  [§](https://www.golang-book.com/books/intro/11#section1)

- In chapter11 folder create another folder called math. Inside of this folder create a file called math.go that contains this:
```go
package math

func Average(xs []float64) float64 {
  total := float64(0)
  for _, value := range xs {
      total += value
  }
  return total / float64(len(xs))
}

```
- use the package like this
```go
package main

import "fmt"
import "golang-book/chapter11/math"

func main() {
  xs := []float64{1,2,3,4}
  avg := math.Average(xs)
  fmt.Println(avg)
}
```

- Using a terminal in the math folder you just created run go install. This will compile the math.go program and create a linkable object file

## Documentation 
- Go has the ability to automatically generate documentation for packages we write in a similar way to the standard package documentation. In a terminal run this command:
```sh
godoc go_learn/golang_book/chapter11/math  Average
```
- also godoc is available via browser

```sh
godoc -http=":6060"
```
- and entering this URL into your browser:

```sh
http://localhost:6060/pkg/
```
You should be able to browse through all of the packages installed on your system.


#Problems
- [x] Why do we use packages?
  - we use packages to:
    1. reuse code
    2. organize code
    3. it's easy to maintain 

- [x] What is the difference between an identifier that starts with a capital letter and one which doesn’t? (Average vs average)
  - with the lower one we limit variables to package only
  - with the capital one we expose the variable or method outside the package

- [x] What is a package alias? How do you make one?
  - the alias is like a variable we associate to a package when we impot it, see example bellow 
   ```go
   import m "go_learn/golang_book/chapter11/math"
   ```

- [x] We copied the average function from chapter 7 to our new package. Create Min and Max functions which find the minimum and maximum values in a slice of float64s.
  - see [solution](./math/math.go)

- [x] How would you document the functions you created in #3?
  with // above the declaration of the function
  ```go
  // Finds the minimum value from the slice xs
  func Min(xs []float64) float64{

  }
  ```