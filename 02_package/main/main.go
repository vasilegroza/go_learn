package main

import (
	"fmt"
	"github.com/vasilegroza/go_learn/02_package/stringutil" // istalled custom package
)

func main() {
	fmt.Println(stringutil.Reverse("!og olleH"))
	fmt.Println(stringutil.MyName)
}
