package main

import "fmt"

func main() {
	fmt.Print("Converting from feet into meter: ")
	var m float64
	var ft float64
	fmt.Scanf("%f", &ft)

	m = (ft) * 0.3048
	fmt.Println(m)
}
