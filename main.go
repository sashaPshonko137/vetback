package main

import "fmt"

func main() {
	var s []int
	a := s[2:2:2]
	_ = a
	fmt.Println(a)
}