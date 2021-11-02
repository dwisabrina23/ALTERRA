package main

import "fmt"

func swap(a, b *int) {
	temp := *b
	*b = *a
	*a = temp
}

func main() {

	a := 10

	b := 20

	swap(&a, &b)

	fmt.Println(a, b)

}
