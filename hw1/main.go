package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const usd = 70.44

func inputFloat(msg string) (input float64) {
	var err error
	fmt.Println(msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	input, err = strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("Вы ввели не число!")
		return
	}
	return input
}

// 1) Func for convert Rub -> USD.
func convertRubToUsd() {
	rub := inputFloat("Введите сумму для конвертации в рублях:")
	if rub == 0 {
		return
	}
	sum := rub / usd
	fmt.Printf("Результат: %.2f USD\nКурс USD: %.2f\n", sum, usd)
}

// 2) Func for count Perimeter , Square, Hypotenuse
func countPSH() {
	legA := inputFloat("Введите катет А прямоугольного треугольика:")
	legB := inputFloat("Введите катет B прямоугольного треугольика:")
	if legA == 0 || legB == 0 {
		return
	}
	sq := legA * legB / 2
	fmt.Printf("Площадь прямоугольного треугольника: %.2f \n", sq)

	hp := math.Sqrt(math.Pow(legA, 2) + math.Pow(legB, 2))
	fmt.Printf("Гипотенуза прямоугольного треугольника: %.2f \n", hp)

	pr := legA + legB + hp
	fmt.Printf("Периметр прямоугольного треугольника: %.2f \n", pr)
}

// 3) Func for Invest Count
func countInvest() {
	sum := inputFloat("Введите сумму вклада:")
	proc := inputFloat("Введите годовой процент банка:")
	if sum == 0 || proc == 0 {
		return
	}
	for i := range [5]int{} {
		sum += sum * proc / 100
		fmt.Printf("Сумма вклада за %d год: %.2f\n", i+1, sum)
	}
}

func main() {
	fmt.Println("1. Написать программу для конвертации рублей в доллары. Программа запрашивает сумму в рублях" +
		"и выводит сумму в долларах. Курс доллара задайте константой.")
	convertRubToUsd()
	fmt.Println("\n2. Даны катеты прямоугольного треугольника. " +
		"Найти его площадь, периметр и гипотенузу. Используйте тип данных float64 и функции из пакета math.")
	countPSH()
	fmt.Println("\n3. Пользователь вводит сумму вклада в банк и годовой процент. Найти сумму вклада через 5 лет.")
	countInvest()
}
