package main

import "fmt"

func PairSum(arr []int, target int) []int {
	result := make(map[int]int)

	for index, num := range arr {
		//buat var untuk bil yg dicari -> hasil pengurangan target dgn num
		temp := target - num
		if val, ok := result[temp]; ok {
			return []int{val, index}
		}
		result[num] = index
	}
	return []int{}

}

func main() {

	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]

}
