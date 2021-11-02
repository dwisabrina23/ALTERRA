package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	var (
		res   []int
		tempA int
		tempB int
	)
	res = append(res, 0)
	fmt.Println("res awal", res)
	for i := 1; i < len(jumps); i++ {
		if i >= 2 {
			fmt.Println(i)
			fmt.Println("jump i", jumps[i], "jump i-1 dan i-2", jumps[i-1], jumps[i-2])
			tempA = res[i-1] + abs(jumps[i]-jumps[i-1])
			tempB = res[i-2] + abs(jumps[i]-jumps[i-2])
			fmt.Println("temp a:", tempA, " temp B:", tempB)
			if tempA < tempB {
				res = append(res, tempA)
				fmt.Println("res setelah if", res)
			} else {
				res = append(res, tempB)
				fmt.Println("res setelah else", res)
			}
		} else {
			res = append(res, abs(jumps[i]-jumps[i-1]))
			fmt.Println("res setelah", res)
		}
	}
	return res[len(res)-1]
}

func abs(num int) int {
	return int(math.Abs(float64(num)))
}

func main() {

	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40

	fmt.Println(Frog([]int{20, 70, 30, 10, 50, 60, 30, 30}))

}
