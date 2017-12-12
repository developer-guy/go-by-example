package main

import (
	"fmt"
	"time"
)

func main() {
	var p = fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)
}
