package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)  // alan adları olmadan yazar
	fmt.Printf("%+v\n", p) // alan adları ile beraber yazar
	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p) // değerin tipini yazar.

	fmt.Printf("%t\n", true)         // bool değeri formatlama
	fmt.Printf("%d\n", 123)          // int değer formatlama
	fmt.Printf("%f\n", 123.32)       // float değer formatlama
	fmt.Printf("%s\n", "\"string\"") // string deger formatlama

	fmt.Printf("%p\n", &p) // pointer değer formatlama

	formattedStr := fmt.Sprintf("a %s", "string") // Sprintf methodu string'i hem formatlar hem geri döner.
	fmt.Println("formattedstr:", formattedStr)

	fmt.Fprintf(os.Stderr, "an %s\n", "error") // Fprintf methodu hem formatlayıp hem yazdırmak için kullanılır.
}
