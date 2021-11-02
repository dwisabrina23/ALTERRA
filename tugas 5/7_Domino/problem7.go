package main

import (
	"fmt"
	"strings"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	target := fmt.Sprintf("%d%d", deck[0], deck[1])
	sum := 0
	result := []int{}
	for _, nums := range cards {
		for _, num := range nums {
			check := fmt.Sprintf("%v", num)
			//cek apakah angkanya ada di deck atau ngga
			if strings.Contains(target, check) {
				if sum < nums[0]+nums[1] {
					sum = nums[0] + nums[1]
					result = append(result, nums[0])
					result = append(result, nums[1])
				}
				break
			}
		}
	}
	//kalau ga ketemu tutup kartu
	if sum == 0 {
		return "tutup kartu"
	}

	return result
}

func main() {

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))

	// [3, 4]

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))

	// [6 5]

	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))

	// “tutup kartu”

}
