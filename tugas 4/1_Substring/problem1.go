package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if strings.Contains(a, b) {
		return b
	}
	return a
}
func main() {
	fmt.Println(Compare("akang", "aka"))
}
