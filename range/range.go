package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}

	sum := 0

	for _, num := range nums {  // indeks,value := range array or map {
		 					 // şeklinde kullanılır.
		sum += num
	}

	fmt.Println("sum:", sum)

	len := len(nums)

	for i := 0; i < len; i ++ {
		fmt.Println("Val : ", nums[i], "Index :", i)
	}

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index of 3 :", i)
		}
	}

	var kvs map[string]string = map[string]string{"a": "apple", "b": "banane"}
	for k, v := range kvs {
		fmt.Printf("Key %s ->  Val %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i,v := range "go" {
		fmt.Println(i,v)
	}

}
