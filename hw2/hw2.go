package main

import (
	"fmt"
	h "github.com/gb_home/hw2/helper"
	prime "github.com/gb_home/hw2/primeNumber"
)

func main() {
	fmt.Println("Написать функцию, которая определяет, четное ли число.")
	fmt.Println("isMod(4):", h.IsMod(4))
	fmt.Println("isMod(7):", h.IsMod(7))
	fmt.Println("Написать функцию, которая определяет, делится ли число без остатка на 3.")
	fmt.Println("isDivide(9):", h.IsDivide(9))
	fmt.Println("isDivide(8):", h.IsDivide(8))
	fmt.Println("Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.")
	h.CreateFib()
	fmt.Println("Заполнить массив из 100 элементов различными простыми числами.")
	prime.CreatePrimeNum()
}
