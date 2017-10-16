package main

import "fmt"

func main() {
	factOfFive := fact(5)
	fmt.Println("factOfFive:", factOfFive)
}

func fact(n int) (int) {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
