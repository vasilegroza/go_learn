package main

import "fmt"

func main() {
	x := 5
	y := 2
	// // var p *int = &x
	// // fmt.Printsln(*p)
	// fmt.Println("x=", x)
	// setZero(&x)
	// fmt.Println("x=", x)
	fmt.Printf("x=%v, y=%v \n", x, y)
	swap(&x, &y)
	fmt.Printf("x=%v, y=%v \n", x, y)

}

func setZero(xPtr *int) {
	fmt.Println("Set 0 at", xPtr)
	*xPtr = 0
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
