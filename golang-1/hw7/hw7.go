//интерфейсы
package main

import (
	"fmt"
	"math"
)

//S = π × r2
type Circle struct {
	radius float64
}

func (S Circle) Area() float64 {
	result := math.Pi * math.Pow(S.radius, 2)
	return result
}

//S = a × b
type Rectangle struct {
	a float64
	b float64
}

func (S Rectangle) Area() float64 {
	result := S.a * S.b
	return result
}

type Result interface {
	Area() float64
}

func main() {
	circle := Circle{
		radius: 100,
	}

	rectangle := Rectangle{
		a: 10,
		b: 20,
	}

	res(circle)
	res(rectangle)
}
func res(a Result) {
	fmt.Println(a.Area())
}
