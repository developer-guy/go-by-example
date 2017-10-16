package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{Name: "Batuhan", Age: 23})

	fmt.Println(person{Name: "Fred"})

	var personVal person = person{Name: "Asena", Age: 20}
	fmt.Println("Person address : ", personVal)

	var copyPerson *person = &personVal

	copyPerson.Name = "Now batuhan"

	fmt.Println(personVal.Name)

}
