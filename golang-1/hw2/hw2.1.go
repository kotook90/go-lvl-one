//Напишите программу для вычисления площади прямоугольника.
//Длины сторон прямоугольника должны вводиться пользователем с клавиатуры
//S=0,5*a*b
package main

import "fmt"

func S(a, b int) int {
	s := a * b / 2
	return s
}
func main() {
	var a int
	var b int
	fmt.Println("Введите длину стороны 'a': ")
	fmt.Scan(&a)
	fmt.Println("Введите длину стороны 'b': ")
	fmt.Scan(&b)
	result := S(a, b)
	fmt.Printf("Площадь прямоугольника равна %d", result)

}
