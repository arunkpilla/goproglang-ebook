package main

import (
	"fmt"
	"time"
)

func main() {
	s := "Hello World."
	var o string
	start := time.Now()
	for i := len(s) - 1; i >= 0; i-- {
		o += string(s[i])
	}
	fmt.Println(string(o))
	fmt.Println(time.Since(start))
}
