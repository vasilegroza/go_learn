package main

import (
	"fmt"
)

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	y := []float64{98, 93, 77, 82, 83}
	total = 0
	for _, value := range y {
		total += value
	}
	// fmt.Println(total / float64(len(y)))

	// ySlice := y[0:2]
	// fmt.Println(ySlice)
	// slice := make([]float64, 3, 9)
	// fmt.Println(len(slice))

	// appendSlice()
	// copySlice()

	doMap()

}

func appendSlice() {
	var slice1 = []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1)
	fmt.Println(slice2)
}

func copySlice() {

	var slice1 = []int{1, 2}
	var slice2 = make([]int, 2)

	copy(slice2, slice1)
	fmt.Println(slice2)

}

func doMap() {
	var x map[string]string
	x = make(map[string]string)
	x["name"] = "Vasile"
	x["nickname"] = "utilizatorvalid"
	nickname, ok := x["nickname"]
	fmt.Println(nickname, ok)
	delete(x, "nickname")
	nickname, ok = x["nickname"]
	fmt.Println(nickname, ok)
}
