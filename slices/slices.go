package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)


	//slice1 := []int{}
	//array1 := [5]int{}

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "new str")

	fmt.Println("set:", s)
	fmt.Println("len:", len(s))

	c := make([]string, len(s))

	copy(c, s)

	fmt.Println("Copy c:", c)

	l := s[2:5]

	fmt.Println("sl1:", l)

	l = s[2:]

	fmt.Println("sl2:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

}
