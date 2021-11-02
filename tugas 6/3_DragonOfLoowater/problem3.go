package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	var i, maxHeight, win = 0, 0, 0
	//urutkan berdasarkan kekuatan biar balance waktu ngelawan
	sort.Ints(knightHeight)
	sort.Ints(dragonHead)

	//kurang pasukan
	if len(knightHeight) < len(dragonHead) {
		fmt.Println("Knight fall")
		return
	}

	//kalau pasukan cukup
	for _, height := range knightHeight {
		if height >= dragonHead[i] {
			win++
			i++ //next dragon
			maxHeight += height
		}
		if i == len(dragonHead) {
			fmt.Println(maxHeight)
			return
		}
	}
	fmt.Println("Knight fall")
}

func main() {

	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4}) // 11

	DragonOfLoowater([]int{5, 10}, []int{5}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10

}
