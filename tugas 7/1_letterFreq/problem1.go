package main

import (
	"fmt"
	"regexp"
	"time"
)

func letterFreq(word string) {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	procesed := reg.ReplaceAllString(word, "")
	countItems := make(map[string]int)

	//hitung jumlah kemunculan tiap elemen
	for _, char := range procesed {
		countItems[string(char)] = countItems[string(char)] + 1
	}

	for key, val := range countItems {
		fmt.Printf("%s : %d\n", key, val)
	}
}
func main() {
	go letterFreq("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
	time.Sleep(3 * time.Second)
}
