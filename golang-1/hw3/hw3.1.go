//Доработать калькулятор: больше операций и валидации данных
package main

import (
	"fmt"
	"math"
	"os"
	//"reflect"
)

func main() {
	var (
		numberOne, numberTwo, result float64
		operation                    string
	)

	fmt.Printf("Добро пожаловать в программу 'Калькулятор'! \n" +
		"Вам доступны следующие операции:\n" +
		"'+' сложение \n" +
		"'-' вычитание \n" +
		"'*' умножение\n" +
		"'/' деление \n" +
		"'sqrt' квадратный корень \n" +
		"'st' возведение в степень\n" +
		"Для получения результата введите первое число, затем, математическую операцию, затем, второе число\n" +
		"Введите первое число: ")
	_, err := fmt.Scan(&numberOne)
	if err != nil {
		fmt.Println("Ошибка.Необходимо ввмести число!", err)
		return
	}

	fmt.Println("Введите математическую операцию: ")
	_, err2 := fmt.Scan(&operation)
	if err2 != nil {
		fmt.Println("Ошибка. Вы ввели неправильную операцию!", err2)
	}

	if operation == "sqrt" {
		result = math.Sqrt(numberOne)
		fmt.Println("Результат: ", result)
		return
	}
	fmt.Println("Введите второе число: ")
	_, err3 := fmt.Scan(&numberTwo)
	if err != nil {
		fmt.Println("Ошибка.Необходимо ввмести число!", err3)
		return
	}

	switch operation {
	case "+":
		result = numberOne + numberTwo

	case "-":
		result = numberOne + numberTwo
	case "*":
		result = numberOne * numberTwo
	case "/":

		result = numberOne / numberTwo
		if numberTwo == 0 {
			fmt.Println("Делить на ноль нельзя!")
			os.Exit(1)
		}

	case "st":
		result = math.Pow(numberOne, numberTwo)
	default:
		fmt.Println("Вы не выбрали операцию!")
	}
	fmt.Println("Результат: ", result)
}
