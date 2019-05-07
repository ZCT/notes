package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["foo"]++
	fmt.Println(m["foo"])
}
