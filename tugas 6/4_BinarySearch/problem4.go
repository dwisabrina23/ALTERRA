package main

import "fmt"

func BinarySearch(array []int, x int) {
	head := 0
	tail := len(array) - 1
	for head <= tail {
		mid := (head + tail) / 2
		if x == array[mid] {
			fmt.Println(mid)
			return
		} else if x < array[mid] {
			tail = mid
		} else {
			head = mid + 1
		}
	}
	fmt.Println("-1")
}

func main() {

	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // 2

	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // 3

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53) // 6

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1

}
