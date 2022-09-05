package main

import "fmt"

func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func noempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
func main() {
	data := []string{"one", "", "three"}
	fmt.Println(noempty(data))
	fmt.Println(data)
	data2 := []string{"one", "", "three"}
	fmt.Println(noempty2(data2))
	fmt.Println(data2)
}
