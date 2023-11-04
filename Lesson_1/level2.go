package Lesson_1

import "math"

func CircleArea(r float64) float64 {
	const pi = math.Pi
	area := pi * r * r
	return area
}
