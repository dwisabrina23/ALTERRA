/* DAY-2
Problem 4
program menentukan apakah bilangan yg diinputkan prima atau tidak*/

package main

import (
	"fmt"
	"math"
)

//fungsi print bilangan prima dengan range n-m
func printPrime(nrange, mrange int) {
	sum := 0
	for i := nrange; i <= mrange; i++ {
		if isPrime(i) {
			fmt.Println(i)
			sum += 1
		}
	}
	fmt.Printf("Jumlah bilangan prima dari %d sampai %d adalah %d", nrange, mrange, sum)
}

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

	printPrime(10, 50)
}
