package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

var dilbert Employee

func EmployeeByID(id int) *Employee {}

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	dilbert.ID = 100
	Name := &dilbert.Name
	*Name = "hogehoge"
	fmt.Printf("%v", *Name)
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Name += "hugahuga"
	fmt.Printf("%v", *Name)
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 20
	w.Spokes = 20
}
