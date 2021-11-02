package main

import (
	"fmt"
)

type student struct {
	name       string
	nameEncode string
	// score int
}

type Chiper interface {
	Encode() string

	Decode() string
}

func (s *student) Encode() string {
	result := []rune{}
	for _, chr := range s.name {
		switch {
		case ('a' <= chr) && (chr <= 'z'):
			// 'z' - chr + 'a'
			result = append(result, 219-chr)
		default:
			continue
		}
	}
	return string(result)
}

func (s *student) Decode() string {
	result := []rune{}
	for _, chr := range s.nameEncode {
		switch {
		case ('a' <= chr) && (chr <= 'z'):
			// 'z' - chr + 'a'
			result = append(result, 219-chr)
		default:
			continue
		}
	}
	return string(result)
}

func main() {

	var menu int

	var a = student{}

	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")

	fmt.Scan(&menu)

	if menu == 1 {

		fmt.Print("\nInput Student’s Name : ")

		fmt.Scan(&a.name)

		fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())

	} else if menu == 2 {

		fmt.Print("\nInput Student’s Encode Name : ")

		fmt.Scan(&a.nameEncode)

		fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())

	} else {

		fmt.Println("Wrong input name menu")

	}

}
