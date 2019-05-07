package main

import "fmt"

func one(s []int) {
	two(s)
	s = append(s, 1)
}

func two(s []int) {
	three(s)
	s = append(s, 2)
}

func three(s []int) {
	s = append(s, 3)
}

func main() {
	s := make([]int, 0, 5)
	one(s)
	fmt.Println(s)
}
