package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""
	var now int64
	now = time.Now().UnixNano()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%v\n", time.Now().UnixNano()-now)
	now = time.Now().UnixNano()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%v\n", time.Now().UnixNano()-now)
}
