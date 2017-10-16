package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// x <- chan string ' ifade sadece channelden değer alınabilir
	// x chan <- string ' ifade sadece channela değer gönderilebilir
	msg := <-pings
	pongs <- msg
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "message received")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
