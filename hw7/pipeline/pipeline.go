package pipeline

import "fmt"

// PipeLine with klimit 10 numbers.
func PipeLine(){
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0;x < 10 ; x++ {
			naturals <- x
		}
		close(naturals)
		close(squares)
	}()

	// возведение в квадрат
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// печать
	for {
		x, ok := <-squares
		if !ok {
			break // канал закрыт и пуст
		}
		fmt.Println(x)
	}

}
