package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.")
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("abort")
		return
	}
	launch()
}
func launch() {
	fmt.Println("launch")
}
