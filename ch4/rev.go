package main

import "fmt"

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(reverse(a[:]))
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s[:])
	fmt.Println(s)
}
