package main

import "fmt"

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}

	sum(nums...)
}
func sum(nums ...int) {
	fmt.Print(nums, "=")
	total := 0

	for i, num := range nums {
		fmt.Println("Index : ", i)
		total += num
	}

	fmt.Print(total, " ")
}
