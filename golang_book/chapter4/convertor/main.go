package main

import (
	"fmt"
)

func main() {
	var temperature_fahrenheit float32
	var temperature_celsius float32
	fmt.Println("Enter the temperature in F")
	fmt.Scanf("%f", &temperature_fahrenheit)
	temperature_celsius = (temperature_fahrenheit - 32) * 5 / 9
	fmt.Println("Temperature in celsius is:", temperature_celsius)

}
