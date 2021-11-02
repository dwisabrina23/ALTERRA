package main

import "fmt"

func moneyCoins(money int) []int {
	coin := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	idx := 0
	var res []int
	for idx < len(coin) {
		if money >= coin[idx] {
			money -= coin[idx]
			res = append(res, coin[idx])
		} else {
			idx++
		}

		if money <= 0 {
			break
		}
	}
	return res
}

func main() {

	fmt.Println(moneyCoins(123)) // [100 20 1 1 1]

	fmt.Println(moneyCoins(432)) // [200 200 20 10 1 1]

	fmt.Println(moneyCoins(543)) // [500, 20, 20, 1, 1, 1]

	fmt.Println(moneyCoins(7752)) // [5000, 2000, 500, 200, 50, 1, 1]

	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]

}
