package main

import (
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	red := color.RGBA{0, 255, 0, 0}
	blue := color.RGBA{0, 0, 0, 255}
	p := ColoredPoint{Point{1, 1}, red}
	q := ColoredPoint{Point{1, 1}, blue}
	p.Distance(q.Point)
}
