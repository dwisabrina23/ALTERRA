// package main

// import "fmt"

// func Solve(arr []int) (res []int) {
// 	visited := map[int]bool{}
// 	for i := len(arr) - 1; i >= 0; i-- {
// 		n := arr[i]
// 		if visited[n] {
// 			continue
// 		}

// 		visited[n] = true
// 		res = append([]int{n}, res...)
// 	}

// 	return
// }

// func main() {
// 	fmt.Println(Solve([]int{2, 3, 4, 3, 3, 3}))
// }
package main

import (
	"fmt"
)

func caesar(step byte, strRaw string) {
	if step > 26 {
		step %= 26
	}
	str_slice_src := []byte(strRaw)
	str_slice_dst := make([]byte, len(str_slice_src), len(str_slice_src))
	for i := 0; i < len(str_slice_src); i++ {
		if str_slice_src[i] < 123-step {
			str_slice_dst[i] = str_slice_src[i] + step
		} else {
			str_slice_dst[i] = str_slice_src[i] + step - 26
		}
	}
	fmt.Println("plain text:", strRaw)
	fmt.Println("ciphertext:", string(str_slice_dst))
}

func main() {

	caesar(3, "abc") // def

	caesar(2, "alta") // cnvc

	caesar(10, "alterraacademy") // kvdobbkkmknowi

	caesar(1, "abcdefghijklmnopqrstuvwxyz")

	// bcdefghijklmnopqrstuvwxyza

	// caesar(1000, "abcdefghijklmnopqrstuvwxyz")

	// mnopqrstuvwxyzabcdefghijkl

}
