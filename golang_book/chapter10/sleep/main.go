package main

import (
	"fmt"
	"time"
)

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	Sleep(5)
	fmt.Println("hello after 5 seconds")
}
