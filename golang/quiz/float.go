package main

import "fmt"

func main() {
	var f1, f2 float64 = 0.1, 0.2
	var f3 float64 = (0.1 + 0.2) / 2
	fmt.Println(f3 == (f1+f2)/2)
}
