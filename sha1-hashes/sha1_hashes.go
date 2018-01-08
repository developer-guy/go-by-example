package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	s := "sha1 this string"

	hash := sha1.New()

	hash.Write([]byte(s))

	bs := hash.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

}
