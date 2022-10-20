package main

import "fmt"

var balance int

func Deposit(amount int) { balance = balance + amount }

func Balance() int { return balance }
func main() {
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
	}()
	go Deposit(100)

	fmt.Println("=", Balance())
}
