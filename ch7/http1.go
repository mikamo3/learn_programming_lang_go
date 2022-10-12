package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollers float64
type database map[string]dollers

func (d dollers) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s:%s\n", item, price)
	}
}
func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
