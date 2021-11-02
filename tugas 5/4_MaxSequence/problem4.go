package main

import (
	"fmt"
)

func tempMaxSequence(arr []int) int {
	max := 0
	tempMax := arr[0]
	for i := 0; i < len(arr); i++ {
		max += arr[i]
		if tempMax < max {
			tempMax = max
		}

		if max < 0 {
			max = 0
		}
	}
	return tempMax
}

func main() {

	fmt.Println(tempMaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	fmt.Println(tempMaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

	fmt.Println(tempMaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

	fmt.Println(tempMaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

	fmt.Println(tempMaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

}
