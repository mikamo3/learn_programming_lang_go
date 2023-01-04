package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
	t := v.Type()
	fmt.Println(t.String())

	x := v.Interface()
	i := x.(int)
	fmt.Printf("%d\n", i)
}
