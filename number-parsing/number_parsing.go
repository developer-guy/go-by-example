package main

import (
	"strconv"
	"fmt"
	"reflect"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(reflect.TypeOf(f))

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(reflect.TypeOf(i))

	i2, _ := strconv.Atoi("135")
	fmt.Println(reflect.TypeOf(i2))
}
