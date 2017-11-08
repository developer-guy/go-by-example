package main

import "fmt"

func main() {
	queue := make(chan string, 2) //buffered channel

	queue <- "one"
	queue <- "two"

	close(queue)
	for element := range queue {
		fmt.Println(element)
	}
}
