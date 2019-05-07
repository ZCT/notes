package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}

	fmt.Println("%v", v)
}

//https://www.do1618.com/archives/1188/go-range-loop-internals/
