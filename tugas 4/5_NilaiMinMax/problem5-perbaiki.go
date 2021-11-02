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
	min = s.score[0]
	for i := 0; i < len(s.score); i++ {
		if s.score[i] < min {
			min = s.score[i]
			name = s.name[i]
		}
	}
	return
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	for i := 0; i < len(s.score); i++ {
		if s.score[i] > max {
			max = s.score[i]
			name = s.name[i]
		}
	}
	return
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
