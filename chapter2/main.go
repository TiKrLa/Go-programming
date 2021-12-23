package main

import "fmt"

func main() {
	fmt.Print("Convert to Celsius: ")
	var celsius float64
	var fahrenfeit float64
	fmt.Scanf("%f", &fahrenfeit)

	celsius = (fahrenfeit - 32) * 5 / 9
	fmt.Println(celsius)
}
