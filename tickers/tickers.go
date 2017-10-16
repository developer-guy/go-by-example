package main

import (
	"time"
	"fmt"
)

func tickAt(ticker *time.Ticker) {
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go tickAt(ticker)

	time.Sleep(time.Millisecond * 2100)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
