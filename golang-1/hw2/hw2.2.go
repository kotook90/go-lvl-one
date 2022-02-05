// Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
//Формула диаметра круга через площадь: D=2√S/π      ДЛИНА ОКРУЖНОСТИ 		l=πd, где π,d — диаметр окружности
package main

import (
	"fmt"
	"math"
)

//диаметр
func diametr(s float64) float64 {
	d := (math.Sqrt(s / math.Pi)) * 2
	return d
}

//длина окружности
func dlina(d float64) float64 {
	l := math.Pi * d
	return l
}

func main() {
	var s float64
	fmt.Println("Введите площадь круга: ")
	fmt.Scan(&s)
	result := diametr(s)
	fmt.Println("Диаметр круга равен: ", result)
	result2 := dlina(result)
	fmt.Println("Длина окружности равна: ", result2)
}
