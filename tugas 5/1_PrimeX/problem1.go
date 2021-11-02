package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	count := 0
	var result int
	for i := 2; count != number; i++ {
		if isPrime(i) {
			count++
			result = i
		}
	}
	return result
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	} else if number == 2 {
		return true
	}
	//jika bil yg dimasukkan lebih dari 2
	var i int = 2
	batas := math.Sqrt(float64(number))
	for i <= int(batas)+1 {
		if number%i == 0 {
			return false
		}
		i++
	}
	return true
}

func main() {

	fmt.Println(primeX(1)) // 2

	fmt.Println(primeX(5)) // 11

	fmt.Println(primeX(8)) // 19

	fmt.Println(primeX(9)) // 23

	fmt.Println(primeX(10)) // 29

}
