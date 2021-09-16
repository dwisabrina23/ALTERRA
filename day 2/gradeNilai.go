/* DAY-2
Problem 2
program menentukan grade nilai siswa*/

package main

import "fmt"

func main() {
	fmt.Println("PROGRAM GRADE NILAI SISWA")

	//inisialisasi variabel
	var nilai int
	var namaSiswa, gradeNilai string

	//input
	fmt.Print("Masukkan nama siswa: ")
	fmt.Scanln(&namaSiswa)
	fmt.Print("Nilai siswa: ")
	fmt.Scanln(&nilai)

	//hitung grade nilai
	if nilai >= 80 && nilai <= 100 {
		gradeNilai = "A"
	} else if nilai >= 65 && nilai <= 79 {
		gradeNilai = "B"
	} else if nilai >= 50 && nilai <= 64 {
		gradeNilai = "C"
	} else if nilai >= 35 && nilai <= 49 {
		gradeNilai = "D"
	} else if nilai >= 0 && nilai <= 34 {
		gradeNilai = "E"
	} else {
		gradeNilai = "Invalid"
	}

	fmt.Println("Hasil capaian siswa===========")
	//output
	fmt.Printf("Nama Siswa: %s", namaSiswa)
	fmt.Printf("\nGrade Nilai: %s", gradeNilai)
}
