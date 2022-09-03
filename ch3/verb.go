package main

import "fmt"

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := 0xdeadbeef
	fmt.Printf("%d %[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q \n", ascii)
	fmt.Printf("%d %[1]c %[1]q \n", unicode)
	fmt.Printf("%d %[1]c %[1]q \n", newline)
}
