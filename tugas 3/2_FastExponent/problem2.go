package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	temp := pow(x, n/2)
	temp *= temp
	if n%2 == 1 {
		temp *= x
	}
	return temp

}

func main() {

	fmt.Println(pow(2, 3)) // 8

	fmt.Println(pow(7, 2)) // 49

	fmt.Println(pow(10, 5)) // 100000

	fmt.Println(pow(17, 6)) // 24137569

	fmt.Println(pow(5, 3)) // 125

}
