package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}
type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
func (p *Point) SlaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	paths := Path{
		{1, 1},
		{2, 5},
		{3, 3},
	}
	paths[0].SlaleBy(2)
	fmt.Println(paths.Distance())
	distance := Point.Distance
	fmt.Println(distance(paths[0], paths[1]))
}

type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
