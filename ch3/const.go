package main

import (
	"fmt"
	"time"
)

func main() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)
	type Weekday int
	const (
		Sun Weekday = iota
		Mon
		Tue
		Wed
		Thu
		Fri
		Sat
	)
	fmt.Println(Sun, Mon, Tue)
	type Flags uint
	const (
		Flag1 Flags = 1 << iota
		Flag2
		Flag3
	)
	fmt.Println(Flag1, Flag2, Flag3)

}
