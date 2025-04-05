package main

import (
	"fmt"
	"time"
)

func main() {
	s := "Hello World."
	ss := []rune(s)
	start := time.Now()
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {

		ss[i], ss[j] = ss[j], ss[i]
	}
	fmt.Println(string(ss))
	fmt.Println(time.Since(start))
}
