/* DAY-2
Problem 3
program mencari faktor bilangan dari bilangan yang diinputkan*/

package main

import "fmt"

func main() {
	fmt.Println("PROGRAM CETAK FAKTOR BILANGAN")
	//input
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bilangan)

	//cari faktor bilangan yg diinput
	fmt.Printf("Faktor dari %d adalah:\n", bilangan)
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}
