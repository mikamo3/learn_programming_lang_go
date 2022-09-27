package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"
	"time"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
func main() {
	fmt.Println(hypot(1, 3))
	f := square
	fmt.Println(f(10))
	fmt.Println(strings.Map(func(r rune) rune { return r + 2 }, "HAL"))
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 3, 5))
}
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding(%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("sever %s failed to respond after %s", url, timeout)
}
