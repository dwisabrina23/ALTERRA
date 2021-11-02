package main

import (
	"fmt"
	"sort"
)

func MostAppearItem(items []string) {
	countItems := make(map[string]int)
	//hitung jumlah kemunculan tiap elemen
	for _, word := range items {
		countItems[word] = countItems[word] + 1
	}

	//sort berdasarkan nilai count di atas
	keys := make([]string, 0, len(countItems))
	for value := range countItems {
		keys = append(keys, value)
	}
	sort.Slice(keys, func(i, j int) bool { return countItems[keys[i]] > countItems[keys[j]] })

	fmt.Println()
	for _, value := range keys {
		fmt.Printf("%s -> %d ", value, countItems[value])
	}
}

func main() {

	MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})

	// golang->1 ruby->2 js->4

	MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})

	// C->1 D->2 B->3 A->4

	MostAppearItem([]string{"football", "basketball", "tenis"})

	// football->1 basketball->1 tenis->1

}
