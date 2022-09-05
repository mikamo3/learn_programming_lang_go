package main

import "fmt"

func main() {
	var ages map[string]int
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)
	ages["hoge"] = 12
	/*
	   if age, ok := ages["hoge"]; !ok {

	   }
	*/
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}

var m = make(map[string]int)

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
