package main

import "fmt"

type GeometricShape interface {
	printArea()
	calculateArea() int
}

type Square struct {
	edge int
}

type Rectangle struct {
	LongEdge  int
	ShortEdge int
}

func (s *Square) printArea() { // yine pointer receiver olarak implemente ettik.
	fmt.Println("Square Edge :", s.edge)
}

func (s *Square) calculateArea() int {
	return 4 * s.edge
}

func (r *Rectangle) printArea() {
	fmt.Println("Long Edge :", r.LongEdge, "Short Edge:", r.ShortEdge)
}

func (r *Rectangle) calculateArea() int {
	return 2 * (r.LongEdge + r.ShortEdge)
}

func printDetailsOfGeometricShape(g GeometricShape) {
	g.printArea()
	fmt.Println("Area :", g.calculateArea())
}

func main() {
	var square *Square = &Square{edge: 10}
	var rectangle *Rectangle = &Rectangle{ShortEdge: 5, LongEdge: 10}

	printDetailsOfGeometricShape(square)
	// interface'i pointer receiver
	// olarak implemente ettiğimiz için struct'ların pointer
	// adreslerini parametre olarak veriyoruz.
	printDetailsOfGeometricShape(rectangle)

}
