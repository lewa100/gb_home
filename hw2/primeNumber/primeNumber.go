package primeNumber

import "fmt"

var(
	primeNum = [100]int{}
	lastPrimeNum = 2
)

//getPrimeNum - func for find prime numbers.
func getPrimeNum() int{
	if primeNum[0] == 0{
		return lastPrimeNum
	}
	nextPrimeNum()

	return lastPrimeNum
}

// nextPrimeNum - func to get the next single number.
func nextPrimeNum(){
	i := 0
	for i < len(primeNum){
		if primeNum[i] == 0 {
			break
		}
		if lastPrimeNum >= primeNum[i]{
			if lastPrimeNum % primeNum[i] == 0{
				lastPrimeNum++
				i = 0
				continue
			}
		}
		i++
	}
}
// CreatePrimeNum create array[100] prime numbers.
func CreatePrimeNum(){
	for i:=0 ; i < len(primeNum);i++{
		primeNum[i] = getPrimeNum()
	}
	fmt.Println(primeNum)
}
