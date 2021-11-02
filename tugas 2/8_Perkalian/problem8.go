/* DAY-2
Problem 8
program mencetak tabel perkalian*/

package main

import (
	"fmt"
)

//fungsi cetak perkalian
func cetakTabelPerkalian(bil int) {
	//loop row
	for i := 1; i <= bil; i++ {
		//loop column
		for j := 1; j <= bil; j++ {
			fmt.Print(i * j)
			fmt.Print("\t")
		}
		fmt.Println("")
	}
}

func main() {
	fmt.Println("Program cetak tabel perkalian")

	//inisialisasi variabel
	var bil int

	//input
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bil)

	//output
	cetakTabelPerkalian(bil)
}
