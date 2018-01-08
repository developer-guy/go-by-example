package main

import "fmt"

type Shape struct {
	Width  int
	Height int
}

func (s Shape) calculateArea() int { //value receiver olarak implemente edilmiş
	return s.Height * s.Width
}

func (s *Shape) printAreaValues() { //pointer receiver olarak impelement edilmiş
	fmt.Println("Height:", s.Height, "Width:", s.Width)
}

func main() {
	shape := Shape{Width: 10, Height: 50}
	shape.printAreaValues()
	fmt.Println("Area:", shape.calculateArea())

	fmt.Println(shape)

}
