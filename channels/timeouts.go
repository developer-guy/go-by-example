package main

import (
	"time"
	"fmt"
)

func main() {

	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case response := <-c1:
		fmt.Println("Response", response)
	case time := <-time.After(time.Second * 1):
		fmt.Println("tiME OUT 1", time)
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result2"
	}()

	select {
	case response2 := <-c2:
		fmt.Println("Response", response2)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout2 ")
	}
}
