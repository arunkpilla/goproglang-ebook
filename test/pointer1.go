package main

import "fmt"

func febno(n *int) int {
	for ne: == 1 {
		*n *= (*n - 1)
		*n--
	}
	n*febno(n-1)
	return *n
}
func main() {
	no:=1
	fmt.Printf("%d\n", febno(&no))

}
