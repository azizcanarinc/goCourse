package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	kısakenar, uzunkenar float64
}

type circle struct {
	yarıçap float64
}

func (r rectangle) area() float64 {
	return r.kısakenar * r.uzunkenar
}
func (c circle) area() float64 {
	return math.Pi * c.yarıçap * c.yarıçap
}

func school(s shape) {
	fmt.Println("seklin alanu : ", s.area())

}

func Demo1() {
	r := rectangle{kısakenar: 10, uzunkenar: 6}
	school(r)

	c := circle{yarıçap: 10}
	school(c)
}
