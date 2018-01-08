package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
The default number generator is deterministic, so it’ll produce the same sequence of numbers each time by default.
To produce varying sequences, give it a seed that changes. Note that this is not safe to use for random numbers you
intend to be secret, use crypto/rand for those.
*/

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(s1)
	fmt.Println(rand.Intn(100))           // 0 <= n < 100
	fmt.Println(rand.Float64())           // 0.0 <= n < 1.0
	fmt.Println((rand.Float64() * 5) + 5) //5.0 <= n < 10.0

}
