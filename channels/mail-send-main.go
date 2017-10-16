package main

import (
	"github.com/developer-guy/go-by-example/mail"
	"fmt"
	"time"
)

func main() {
	doneChannel := make(chan bool)

	go func() {
		mail.SendMail(doneChannel)
	}()

	fmt.Println("Sending...")
	select {
	case result := <-doneChannel:
		fmt.Println("Send status :", result)
	case <-time.After(time.Second * 1):
		fmt.Println("Mail send timeout")
	}

}
