package helper

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const usd = 70.44
// ConvertRubToUsd for convert Rub -> USD.
func ConvertRubToUsd() {
	rub := inputFloat("Введите сумму для конвертации в рублях:")
	if rub == 0 {
		return
	}
	sum := rub / usd
	fmt.Printf("Результат: %.2f USD\nКурс USD: %.2f\n", sum, usd)
}

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

// CountPSH for count Perimeter , Square, Hypotenuse
func CountPSH() {
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

// CountInvest for Invest Count
func CountInvest() {
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
