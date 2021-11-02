/* DAY-2
Problem 5
program menentukan pangkat dari bilangan

ketentuan:
base (+) pangkat (+) 		= hasil (+)			-> 2^3 = 8
base (+) pangkat (-) 		= 1/hasil (+)		-> 2^-3 = 1/8
base (-) pangkat ganjil (+) = hasil (-)			-> -2^3 = -8
base (-) pangkat ganjil (-) = 1/hasil(-)		-> -2^-3 = -1/8
base (-) pangkat genap  (+) = hasil (+)			-> -2^4 = 16
base (-) pangkat genap  (-) = 1/hasil(+)		-> -2^-4 = 1/16
*/

package main

import (
	"fmt"
	"math"
)

//fungsi kalkulasi pangkat
func hasilPangkat(bil, pangkat int) int {
	//inisiasi pangkat untuk 0 -> semua bil yg dipangkatkan 0 akan menghasilkan 1
	hasil := 1
	for pangkat != 0 {
		hasil = hasil * bil
		pangkat--
	}
	return hasil
}

func main() {
	fmt.Println("PROGRAM MENENTUKAN PALINDROME")

	//inisialisasi variabel
	var base, exponent, result int

	//input
	fmt.Print("Masukkan bilangan pokok: ")
	fmt.Scanln(&base)
	fmt.Print("Masukkan bilangan pangkat: ")
	fmt.Scanln(&exponent)

	//kalkulasi
	pangkatAbs := math.Abs(float64(exponent))
	result = hasilPangkat(base, int(pangkatAbs))

	//output
	//base positif----------------------------
	if base >= 0 {
		//exponen (+)
		if exponent >= 0 {
			fmt.Printf("Hasil dari %d^%d = %d", base, exponent, result)
		} else { //eksponen (-)
			// result = 1/result
			fmt.Printf("Hasil dari %d^%d = 1/%d", base, exponent, result)
		}

	} else { //base negatif---------------------
		if exponent%2 == 1 { //pangkat ganjil
			if exponent >= 0 {
				// result = -(result)
				fmt.Printf("Hasil dari %d^%d = %d", base, exponent, result)
			} else {
				// result = -1/result
				fmt.Printf("Hasil dari %d^%d = -1/%d", base, exponent, result)
			}
		} else { //pangkat genap
			if exponent >= 0 {
				fmt.Printf("Hasil dari %d^%d = %d", base, exponent, result)
			} else {
				// result = 1/result
				fmt.Printf("Hasil dari %d^%d = 1/%d", base, exponent, result)
			}

		}
	}
}
