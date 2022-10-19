package main

import "os"

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
}
