package main

import (
	"fmt"
	"time"
)

func work(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for w := 1; w <= 3; w ++ {
		go work(w, jobs, result)
	}

	for j := 1; j <= 5; j ++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 5; a ++ {
		<-result
	}
}
