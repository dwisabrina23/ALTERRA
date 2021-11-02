package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	var result = []byte{}

	//kalau shit nya lebih dari jumlah abjad
	if offset > 26 {
		offset %= 26
	}

	for _, val := range input {
		//tambah dec val dengan offset
		val += rune(offset)
		temp := val
		if val > 122 {
			temp = 96 + (val % 122)
		}
		result = append(result, byte(temp))
	}
	return string(result)
}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // cnvc

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))

	// bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

	// mnopqrstuvwxyzabcdefghijkl

}
