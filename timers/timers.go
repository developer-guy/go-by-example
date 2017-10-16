package main

import (
	"fmt"
	"time"
)

func main() {
	var timer *time.Timer = time.NewTimer(time.Second * 2)
	fmt.Println("Time now :", time.Now())

	<-timer.C
	fmt.Println("Time now :", time.Now())

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired")
	}()

	var isStop bool = timer2.Stop()
	if isStop {
		fmt.Println("Timer 2 stopped", time.Now())
	}
}
