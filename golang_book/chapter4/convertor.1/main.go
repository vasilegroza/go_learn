package main

import (
	"fmt"
)

func main() {
	var distance_feet float32
	var distance_meters float32
	fmt.Println("Enter the distance in ft")
	fmt.Scanf("%f", &distance_feet)
	distance_meters = distance_feet * 0.3048
	fmt.Println("Distance meters is:", distance_meters)

}
