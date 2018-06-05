package main

import (
	"fmt"
)

func main() {

	add := func(x int, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
	fmt.Println(average([]float64{1, 2, 3}))
	defer fmt.Println("1+2+3+4 = ", sum(1, 2, 3, 4))
	defer fmt.Println("5! = ", factorial(5))
	defer fmt.Println("fib(5) = ", fib(5))
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println("max ", getMax(1, 9, 34, 51, 0, 2))

}
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (rez uint) {
		rez = i
		i += 2
		return rez
	}
}
func average(elements []float64) float64 {
	total := 0.0
	for _, value := range elements {
		total += value
	}
	return total / float64(len(elements))
}
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)

}

func sum(args ...int) (sum int) {
	sum = 0
	for _, value := range args {
		sum += value
	}
	return
}

func half(number int) bool {
	if number%2 == 1 {
		return false
	}
	return false

}

func getMax(args ...int) (max int) {
	if len(args) == 0 {
		panic("zero elements")
	}
	max = args[0]
	for _, value := range args {
		if value > max {
			max = value
		}
	}
	return max
}
