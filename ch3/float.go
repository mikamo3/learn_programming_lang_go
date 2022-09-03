package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxFloat64 == math.MaxFloat64-1)

	for x := 0; x < 8; x++ {
		fmt.Printf("x=%d ex=%8.3f\n", x, math.Exp(float64(x)))
	}
}
