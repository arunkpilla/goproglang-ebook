package main

import "fmt"

func main() {
	s := "Hello Arun"
	for i, v := range s {
		fmt.Printf("Index:%d,Value:%q\n", i, v)
	}
}
