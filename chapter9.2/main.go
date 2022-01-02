package main

import "fmt"

func greatest(args ...int) int {
	max := args[0]
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}
func main() {
	fmt.Println(greatest(12, 94, 128, 457, 8, 50))
}
