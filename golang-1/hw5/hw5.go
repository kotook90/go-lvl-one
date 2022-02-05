//1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
//2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.
package main

import (
	"fmt"
	"os"
)

func main() {
	var number uint
	var exit string
	mapResult := make(map[uint]uint)
	err1 := fmt.Errorf("Вы не ввели число или ввели символ: ")
	err2 := fmt.Errorf("Вы не выбрали y или n, либо ввели число, работа програмы завершена")
	fmt.Println("Введите число и получите его номер в ряду Фибоначчи!(Число должно быть больше нуля)")
	for {
		if _, err := fmt.Scanln(&number); err != nil {
			fmt.Println(err1, err)
			os.Exit(1)
		}
		result := fibbonachi(number)

		mapResult[number] = result
		fmt.Printf("Для числа %d порядковый номер в ряду Фибоначчи равен %d\n", number, result)
		fmt.Printf("Желаете продолжить? y/n\n")
		if _, err := fmt.Scanln(&exit); err != nil {
			fmt.Println(err2, err)

			os.Exit(1)
		}
		if exit == "y" {
			fmt.Printf("Введите следующее число: \n")
			continue
		} else if exit == "n" {

			fmt.Printf("Тут мы сохранили для Вас историю вычислений:\n")
			for i, v := range mapResult {
				fmt.Printf("Для числа %d результат равен %d\n", i, v)
			}
			fmt.Printf("Работа прогрммы звершена, спасибо!\n")
			break
		} else {
			fmt.Println(err2)
			break
		}
	}
}

func fibbonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}
