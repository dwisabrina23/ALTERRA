package main

import (
	"fmt"
	"math"
)

func primaSegiEmpat(high, wide, start int) {
	sum := 0
	start++
	var i, j int = 0, 0
	for i < wide {
		if isPrime(start) {
			fmt.Printf("%d\t", start)
			sum += start
			//geser kanan buat col baru
			j++
		}

		if high == j {
			fmt.Println("")
			//reset j mulai dari col 0, tambahkan i utk row baru
			j = 0
			i++
		}

		start++
	}
	fmt.Println(sum)
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

	primaSegiEmpat(2, 3, 13)

	/*

	   17 19

	   23 29

	   31 37

	   156

	*/

	primaSegiEmpat(5, 2, 1)

	/*

	   2  3  5  7 11

	   13 17 19 23 29

	   129

	*/

}
