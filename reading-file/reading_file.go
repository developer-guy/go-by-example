package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	data, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(data))

	f, err := os.Open("/tmp/dat")
	check(err)

	defer f.Close()

	b1 := make([]byte, len(data))
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

}
