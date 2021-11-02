/* DAY-2
Problem 4
program menentukan apakah bilangan yg diinputkan prima atau tidak*/

package main

import (
	"fmt"
	"math"
)

//fungsi cek bilangan prima
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
	fmt.Println("PROGRAM MENENTUKAN BILANGAN PRIMA")

	//inisialisasi variabel
	var bilangan int

	//input
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bilangan)

	//output
	if isPrime(bilangan) {
		fmt.Println("Prima")
	} else {
		fmt.Println("Bukan Prima")
	}
}
