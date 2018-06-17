package main

import (
	"fmt"
	m "go_learn/golang_book/chapter11/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	min := m.Min(xs)
	max := m.Max(xs)

	fmt.Println("avg=", avg)
	fmt.Println("min=", min)
	fmt.Println("max=", max)

}
