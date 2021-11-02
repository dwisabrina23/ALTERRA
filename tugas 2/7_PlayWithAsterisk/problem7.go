/* DAY-2
Problem 7
program mencetak segitiga asterik berdasarkan bilangan*/

package main

import (
	"fmt"
)

//fungsi cetak segitiga
func playWithAsterik(bil int) {
	for i := 1; i <= bil; i++ {
		for j := bil - i; j > 0; j-- {
			fmt.Print(" ")
		}
		for k := 1; k < i; k++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

func main() {
	fmt.Println("PROGRAM SEGITIGIA ASTERIK")

	//inisialisasi variabel
	var bil int

	//input
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bil)

	//output
	playWithAsterik(bil)
}
