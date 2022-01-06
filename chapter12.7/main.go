package main

import "fmt"

func main() {
	c := make(chan string, 20)
	c <- "Capacity of 20"
	c <- "Hello Go!"
	fmt.Println(<-c)
	fmt.Println(<-c)
}
