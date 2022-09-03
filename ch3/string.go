package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello 世界"
	x := `a
	b
	c`
	fmt.Println(len(s))
	fmt.Println(x)
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println(utf8.RuneCountInString(s))
	ss := "プログラム"
	fmt.Printf("%x\n", ss)
	r := []rune(ss)
	fmt.Printf("%x\n", r)

}
