package main

import (
	"fmt"
	"time"
)

func main() {
	//creating channel using make keyword
	c := make(chan string, 2)

	//calling Sleep function of go
	go func() {
		for {
			c <- "Sleeping after"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c:
				fmt.Println("Message 1", msg1)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
