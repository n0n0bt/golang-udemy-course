package main

import(
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64  {
	return 0.5 * t.base *t.height
}

func (s square) getArea() float64  {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}


func main()  {
	tr := triangle{base: 7.2, height: 3.5}
	sq := square{sideLength: 5.0}

	printArea(tr)
	printArea(sq)
}

