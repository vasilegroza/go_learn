package main

import "fmt"

func main() {
	fmt.Println("Enter a number")
	var input int
	fmt.Scanf("%d", &input)
	var output int = input * 2
	fmt.Println(output)
}

func greetings() {
	var message string = "Hi there"
	var name string
	name = "human"
	welcome := "Welcome to our world"

	fmt.Println(message, name, "!")
	fmt.Println(welcome)
}
