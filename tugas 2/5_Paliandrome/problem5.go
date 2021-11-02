/* DAY-2
Problem 5
program menentukan apakah string yg diinputkan palindrome atau tidak*/

package main

import "fmt"

//fungsi cek palindrome
func isPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("PROGRAM MENENTUKAN PALINDROME")

	//inisialisasi variabel
	var word, result string

	//input
	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&word)

	//output
	if isPalindrome(word) {
		result = "Palindrome"
	} else {
		result = "bukan Palindrome"
	}
	fmt.Printf("%s merupakan %s", word, result)

}
