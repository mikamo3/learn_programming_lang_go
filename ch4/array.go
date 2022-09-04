package main

import "fmt"

func main() {

	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Println(i, v)
	}
	q := [...]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}

	for i, v := range q {
		fmt.Println(i, v)
	}
	for i, v := range r {
		fmt.Println(i, v)
	}
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "u", EUR: "e", GBP: "g", RMB: "r"}
	fmt.Println(RMB, symbol[RMB])
	rr := [...]int{99: -1}
	fmt.Println(rr)
}
