package main

import (
	"io/ioutil"
	"os"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	data := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/data1", data, 0644)
	check(err)

	f, err := os.Create("/tmp/data2")
	check(err)

	helloStr := "Hello my name is Batuhan"
	wroteByteSize, err := f.Write([]byte(helloStr))
	check(err)
	fmt.Printf("wrote %d bytes\n", wroteByteSize)

	defer f.Close()
}
