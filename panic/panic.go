package main

import (
	"os"
	"fmt"
)

func main() {
	panic("a problem")

	if file, err := os.Create("/tmp/file"); err != nil {
		panic(err)
	}else{
		fmt.Println(file.Name())
	}

}
