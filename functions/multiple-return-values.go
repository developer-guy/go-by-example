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
	x,t = swap2(i,y)
	return
}

func swap2(i,y int) (int,int){
	return y,i
}
