package main

import "fmt"

func anag(str1, str2 string) string {
	ss := []rune(str1)
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}

	if string(ss) == string(str2) {
		return (fmt.Sprintf("%s and %s are anagrams", str1, str2))
	} else {
		return (fmt.Sprintf("%s and %s are NOT anagrams", str1, str2))
	}
}
func main() {
	fmt.Println(anag("Hello", "olleH"))
	fmt.Println(anag("Arun", "nurA"))
}
