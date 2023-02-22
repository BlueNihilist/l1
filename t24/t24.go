package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(xVal, yVal float64) Point {
	p := Point{xVal, yVal}
	return p
}

func getDistance(p1, p2 Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func main() {
	point_1 := newPoint(1.0, 2.0)
	point_2 := newPoint(0.0, -0.125)

	fmt.Printf("Point 1: %v\n", point_1)
	fmt.Printf("Point 2: %v\n", point_2)

	fmt.Printf("Distance between them: %f\n", getDistance(point_1, point_2))
}
