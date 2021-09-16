/* DAY-2
Problem 1
program menghitung luas permukaan Tabung*/

package main

import "fmt"

func main() {
	//inisialisasi variabel
	var luasPermukaan, T, r float64
	const pi float64 = 3.14

	fmt.Println("PROGRAM MENGHITUNG LUAS PERMUKAAN TABUNG")
	//input
	fmt.Print("Masukkan tinggi Tabung: ")
	fmt.Scanln(&T)
	fmt.Print("Masukkan jari-jari Tabung: ")
	fmt.Scanln(&r)

	//kalkulasi Luas Permukaan
	luasPermukaan = 2 * pi * r * (r + T)

	//tampilkan hasilnya
	fmt.Println("Luas permukaan tabung: ", luasPermukaan)
}
