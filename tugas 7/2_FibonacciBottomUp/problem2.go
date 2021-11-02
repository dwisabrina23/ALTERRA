package main

import "fmt"

func fibo(n int) int {
	var res, a, b int
	a, b = 0, 1
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		for i := 2; i <= n; i++ {
			res = a + b
			a = b
			b = res
		}
	}
	return res
}

func main() {

	fmt.Println(fibo(0)) // 0

	fmt.Println(fibo(1)) // 1

	fmt.Println(fibo(2)) // 1

	fmt.Println(fibo(3)) // 2

	fmt.Println(fibo(5)) // 5

	fmt.Println(fibo(6)) // 8

	fmt.Println(fibo(7)) // 13

	fmt.Println(fibo(9)) // 34

	fmt.Println(fibo(10)) // 55

}
