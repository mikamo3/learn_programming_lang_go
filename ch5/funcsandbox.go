package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
func main() {
	fmt.Println(hypot(1, 3))
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
