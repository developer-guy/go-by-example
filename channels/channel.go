package main

import "fmt"

func main() {

	messages := make(chan string)
	// channel'a değer göndermek için
	// chan <- sentaksı kullanılır değer almak için ise
	// <- chan sentaksı kullanılır.

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}
