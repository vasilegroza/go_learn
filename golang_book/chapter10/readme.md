# Concurrency [§](https://www.golang-book.com/books/intro/9)

- Go has rich support for concurrency using goroutines and channels.

## Goroutines [§](https://www.golang-book.com/books/intro/9#section1)

- A goroutine is a function that is capable of running concurrently with other functions. To create a goroutine we use the keyword go followed by a function invocation
- see [example](./main.go)

## Channels [§](https://www.golang-book.com/books/intro/9#section2)

- Channels provide a way for two goroutines to communicate with one another and synchronize their execution. Here is an example program using channels:

- see [example](./channels/main.go)
- A channel type is represented with the keyword chan followed by the type of the things that are passed on the channel (in this case we are passing strings). 
- The **<-** (left arrow) operator is used to send and receive messages on the channel.**c <- "ping"** means send "ping". **msg := <- c** means receive a message and store it in msg.

### Channel Direction
- We can specify a direction on a channel type thus restricting it to either sending or receiving. For example pinger's function signature can be changed to this:
 ```go
 func pinger(c chan<- string)
 ```
 - Now c can only be sent to. Attempting to receive from c will result in a compiler error. Similarly we can change printer to this:
  ```go
  func printer(c <-chan string)
  ```
 - A channel that doesn't have these restrictions is known as bi-directional.

### Select
 - Go has a special statement called select which works like a switch but for channels:
 - 
    ```go
    func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
      for {
        c1 <- "from 1"
        time.Sleep(time.Second * 2)
        }
    }()

    go func() {
      for {
        c2 <- "from 2"
        time.Sleep(time.Second * 3)
        }
    }()

    go func() {
      for {
        select {
        case msg1 := <- c1:
            fmt.Println(msg1)
        case msg2 := <- c2:
          fmt.Println(msg2)
        }
    }
    }()

    var input string
    fmt.Scanln(&input)
    ```
### Buffered Channels
```go
c := make(chan int, 1)
```

# Problems
- [x] How do you specify the direction of a channel type?
    - in the definition of the function use <- operator
        - c <-chan string //read direction channel
        - c chan<- string //write direction channel

- [x] Write your own Sleep function using time.After.
    - see [solution](./sleep/main.go)

- [x] What is a buffered channel? How would you create one with a capacity of 20?
    - a buffered channel is a channel that threats messages in asynchronous manner 
    - make(chan int, 20)  