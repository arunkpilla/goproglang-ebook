package main

import "fmt"

var a int

func fb() *int {
	b := 1
	return &b
}
func fa() *int {
	a = 2
	return &a
}
func main() {
	fmt.Printf("%d && %d\n", fb(), fb())
	fmt.Printf("%d && %d", fa(), fa())
}
