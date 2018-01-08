package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"sync/atomic"
)

func main() {
	var ops uint64 = 0
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		atomic.AddUint64(&ops, 1)
		sig := <-signals
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting after ", atomic.LoadUint64(&ops), " times run ")

}
