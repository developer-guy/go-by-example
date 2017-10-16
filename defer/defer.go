package main

import (
	"os"
	"fmt"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(s string) *os.File {
	fmt.Println("Creating")
	file, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(file *os.File) {
	fmt.Println("wiriting")
	fmt.Fprintln(file, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
