package transform

import (
	"github.com/Camburgaler/scholar-utils/pkg/output"
)

// Point represents a coordinate in 2D space.
type Point struct {
	X, Y float64
}

// calculateSlopeIntercept computes the slope and y-intercept from two points.
func calculateSlopeIntercept(p1, p2 Point) output.SlopeIntercept {
	// Handle vertical line case to avoid division by zero
	if p2.X == p1.X {
		return output.SlopeIntercept{Slope: 0, Intercept: p1.Y} // or handle as needed
	}

	m := (p2.Y - p1.Y) / (p2.X - p1.X)
	b := p1.Y - m*p1.X

	return output.SlopeIntercept{Slope: m, Intercept: b}
}
