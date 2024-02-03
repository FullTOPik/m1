package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func CreatePoint(x float64, y float64) Point {
	return Point{X: x, Y: y}
}

func (p *Point) CalculateDistance(point Point) float64 {
	distanceX := p.X - point.X
	distanceY := p.Y - point.Y

	return math.Sqrt(distanceX*distanceX + distanceY*distanceY)
}

func main() {
	fmt.Printf("Enter a first point:\nX: ")

	var x, y float64
	fmt.Scan(&x)
	fmt.Printf("Y: ")
	fmt.Scan(&y)

	firstPoint := CreatePoint(x, y)

	fmt.Printf("\n\nEnter a second point:\nX: ")

	fmt.Scan(&x)
	fmt.Printf("Y: ")
	fmt.Scan(&y)

	secondPoint := CreatePoint(x, y)

	fmt.Printf("\n\nCalculated distance: %f\n", firstPoint.CalculateDistance(secondPoint))
}
