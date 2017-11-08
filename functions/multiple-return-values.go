package main

import "fmt"

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	x,y := swap(5,10)
	fmt.Println("Swapped:",x,y)
}

func vals() (int, int) {
	return 3, 7
}

func swap(i, y int) (x int, t int) {
	x = y
	t = i
	return
}
