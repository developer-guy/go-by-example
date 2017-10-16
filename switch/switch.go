package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend", time.Now().Weekday())
	default:
		fmt.Println("It's a weekday", time.Now().Weekday())
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's a before noon")
	default:
		fmt.Println("It's a after noon")
	}

	findMyType := func(i interface{}) {
		switch fieldType := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", fieldType)
		}
	}

	findMyType(true)
	findMyType(1)
	findMyType("string")
}
