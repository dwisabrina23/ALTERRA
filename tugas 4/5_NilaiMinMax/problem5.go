package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	var sum = 0
	for _, val := range s.score {
		sum += val
	}
	return float64(sum) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	merge := map[string]int{}
	for _, x := range s.name {
		for _, y := range s.score {
			merge[x] = y
		}
	}

	//cari score terendah
	for key, val := range s.score {
		switch {
		case key == 0:
			min = val
		case val < min:
			min = val
		}
	}

	//cari nama
	for key, score := range merge {
		if score == min {
			name = key
		}
	}

	return min, string(name)
}

func (s Student) Max() (max int, name string) {
	merge := map[string]int{}
	for _, x := range s.name {
		for _, y := range s.score {
			merge[x] = y
		}
	}

	//cari score terendah
	for key, val := range s.score {
		switch {
		case key == 0:
			max = val
		case val > max:
			max = val
		}
	}

	//cari nama
	for key, score := range merge {
		if score == max {
			name = key
		}
	}
	return max, string(name)

}

func main() {

	var a = Student{}

	for i := 1; i < 6; i++ {

		var name string

		fmt.Printf("Input %d Student’s Name : ", i)

		fmt.Scan(&name)

		a.name = append(a.name, name)

		var score int

		fmt.Print("Input " + name + " Score : ")

		fmt.Scan(&score)

		a.score = append(a.score, score)

	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())

	scoreMax, nameMax := a.Max()

	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()

	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")

}
