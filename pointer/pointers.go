package main

import "fmt"

func main() {
	i := 5
	zeroval(i)
	fmt.Println("i", i)

	zeroptr(&i) // &i sentaksı i değişkeninin bellek adresini döner
	fmt.Println("i", i)
}

func zeroval(i int) { // i'nin kopyası iletilir.
	i = 0
}

func zeroptr(i *int) { // i değişkenin tutulduğu bellek adresi
	fmt.Println("Addres of i ", i)
	*i = 0 // *i sentaksı ise i değişkenin bulunduğu adresten değerini döndürür.
}
