package main

import "fmt"

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
}

func vals() (int, int) {
	return 3, 7
}
