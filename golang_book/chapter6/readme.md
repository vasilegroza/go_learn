# Arrays, Slices and Maps [§](https://www.golang-book.com/books/intro/6)

## Arrays [§](https://www.golang-book.com/books/intro/6#section1)

  - An array is a numbered sequence of elements of a single type with a fixed length. In Go they look like this:
        ```go
        var x[5]int
        ```
 - **x** is an example of an array which is composed of 5 **int**s.
 + Go also provides a shorter syntax for creating arrays:
    ```go
    x := [5]float64{ 98, 93, 77, 82, 83 }
    ```

## Slices  [§](https://www.golang-book.com/books/intro/6#section2)
- A slice is a segment of an array. Like arrays slices are indexable and have a length. Unlike arrays this length is allowed to change. Here's an example of a slice:
- Here's an example of slice
    ```go
    var x []float64
    ```
- If you want to create a slice you should use the built-in make function:
    ```go
    x := make([]float64, 5)
    ```
- Another way to create slices is to use the [low : high] expression:
    ```go
    arr := [5]float64{1,2,3,4,5}
    x := arr[0:5]
    ```
### Slice Functions
- Go includes two built-in functions to assist with slices: **append** and **copy**
    - append
        ```go
        slice1 := []int{1,2,3}
        slice2 = append(slice1, 4, 5)
        ```
    - copy
        ```go
        slice1 = []int{1,2}
        slice2 = make([]int, 2)
        copy(slice2, slice1)
        ```
## Maps [§](https://www.golang-book.com/books/intro/6#section3)
- A map is an unordered collection of key-value pairs. Also known as an associative array, a hash table or a dictionary, maps are used to look up a value by its associated key.

    ```go
    var x map[string]int
    ```
- The map type is represented by the keyword **map**, followed by the **key type** in brackets and finally the **value type**.

-  Maps have to be initialized before they can be used
    ```go
    var x map[string]int
    x = make(map[stirng]int)
    x["grade"] = 10
    ```
- We can also delete items from a map using the built-in delete function:
    ```go
    delete(x,1)
    ```
- There may not be a key in the map then it will return type zero value. We could check for zero value. Go provides a better way
    ```go
    x["name"] = "Yami"
	x["nickname"] = "bull"
	nickname, ok := x["nickname"]
    ```
    - The first value is the result of the lookup, the second tells us whether or not the lookup was successful.
    ```go
     if nickname, ok := x[nickname]; ok {
         fmt.Println(nickname)
     }
    ```

# Problems
- [x] How do you access the 4th element of an array or slice?
    - x[5]

- [x] What is the length of a slice created using: make([]int, 3, 9)?
    - 3
- [x] Given the array:

    ```go
    x := [6]string{"a","b","c","d","e","f"}
    ```
    - what would x[2:5] give you?
    - [c, d, e]
- [x] Write a program that finds the smallest number in this list:
    ```go
    x := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }
    ```
    - see [solution](./min/main.go)
