package helper

import "fmt"

var fib = [100]int64{}

// IsMod - func output true if mod(number) == 0 ,else - false.
func IsMod(num int) bool{
	if num % 2 == 0{
		return true
	}
	return false
}

// IsDivide - func output true if num divide without remainder(/3).
func IsDivide(numbs ...int) bool{
	var divide int
	if len(numbs) == 1{
		divide = 3
	} else{
		divide = numbs[1]
	}

	if numbs[0] % divide == 0{
		return true
	}
	return false
}

// fibonachi - func for count fibonachi value.
func fibonachi(n int64)int64{
	if n < 2 {
		return n
	}
	return fib[n-1] + fib[n-2]
}

// CreateFib - func for fill Fibonachi array.
func CreateFib(){
	var i int64
	for i=0;i<100;i++ {
		fib[i] = fibonachi(i)
	}
	fmt.Println(fib)
}

