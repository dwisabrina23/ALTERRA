package main

import (
	"fmt"
	"time"
)

func hello(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(i, " ")
		fmt.Println(s)
	}
}

func main() {
	go hello("world")
	hello("hello")
	fmt.Println("main func")
	time.Sleep(1 * time.Second)
}
