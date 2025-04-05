package main

import "fmt"

func summ(s string) uint8 {
	var count uint8
	for i := 0; i < len(s); i++ {
		count += s[i]
	}
	return count
}
func anag(str1, str2 string) string {
	str1_count := summ(str1)
	str2_count := summ(str2)

	if len(str1) == len(str2) && str1_count == str2_count {
		return (fmt.Sprintf("%s and %s are anagrams", str1, str2))
	} else {
		return (fmt.Sprintf("%s and %s are NOT anagrams", str1, str2))
	}
}
func main() {
	fmt.Println(anag("Hello", "olleH"))
	fmt.Println(anag("apple", "pale"))
	fmt.Println(anag("night", "thing"))
	fmt.Println(anag("night", "Thing"))
	fmt.Println(anag("race", "care"))
}
