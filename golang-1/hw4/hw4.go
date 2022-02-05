//Познакомьтесь с алгоритмом сортировки вставками. Напишите приложение, которое принимает
//на вход набор целых чисел и отдаёт на выходе его же в отсортированном виде.
//Сохраните код, он понадобится нам в дальнейших уроках.
package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	var exit string
	sortSlice := []int{}
	err1 := fmt.Errorf("Вы не ввели число или ввели символ: ")
	err2 := fmt.Errorf("Вы не выбрали y или n, либо ввели число, работа програмы завершена")
	fmt.Println("Введите последовательно несколько чисел и программа отсортирует их от меньшего к большему")
	for {
		if _, err := fmt.Scanln(&number); err != nil {
			fmt.Println(err1, err)
			os.Exit(1)
		}
		sortSlice = append(sortSlice, number)

		fmt.Printf("Желаете продолжить? y/n\n")
		if _, err := fmt.Scanln(&exit); err != nil {
			fmt.Println(err2, err)
			os.Exit(1)
		}
		if exit == "y" {
			fmt.Printf("Введите следующее число: \n")
			continue
		} else if exit == "n" {
			InsertionSort(sortSlice)
			fmt.Println(sortSlice)
			fmt.Printf("Работа прогрммы звершена, спасибо!\n")
			break
		} else {
			fmt.Println(err2)
			break
		}
	}
}

func InsertionSort(ar []int) {

	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}

}
