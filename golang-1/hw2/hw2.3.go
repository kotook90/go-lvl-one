//С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе
package main

import "fmt"

func main() {
	var (
		a, sot, des, ed int
	)
	fmt.Println("Введите число: ")
	fmt.Scan(&a)
	sot = a / 100
	des = a / 10
	ed = a / 1
	fmt.Printf("В Вашем числе %d сотен, %d десятков и %d единиц", sot, des, ed)
}
