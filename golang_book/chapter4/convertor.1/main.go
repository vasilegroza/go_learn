package main

import (
	"fmt"
)

func main() {
	var distanceFeet float32
	var distanceMeters float32
	fmt.Println("Enter the distance in ft")
	fmt.Scanf("%f", &distanceFeet)
	distanceMeters = distanceFeet * 0.3048
	fmt.Println("Distance meters is:", distanceMeters)

}
