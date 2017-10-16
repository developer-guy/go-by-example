package main

import (
	"errors"
	"fmt"
	"strconv"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant work with 42")
	}

	return arg + 3, nil
}

type CustomError struct {
	arg  int
	prob string
}

func (c *CustomError) Error() string { //pointer receiver olarak error interfacesinin Error methodunu implemente ettik.
	return fmt.Sprintf("%d - %s", c.arg, c.prob)
}

func f2(arg int) (int, *CustomError) {
	if arg == 42 {
		return -1, &CustomError{arg, "cant work with " + strconv.Itoa(arg)}
	}
	return arg + 3, nil //burda geriye customError'ün pointerı geriye döndürüldüğü için nil geri döndürebildik.
}

func main() {

	arg, error := f1(42)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Arg:", arg)
	}


	var arg2 int
	var customError *CustomError
	arg2, customError = f2(42)

	if customError != nil {
		fmt.Println("Error explanation:",customError.prob)
	} else {
		fmt.Println("Arg:", arg2)
	}
}
