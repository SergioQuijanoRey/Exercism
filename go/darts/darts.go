// Package darts has the function Score which calculates the score of a given shot
package darts

import "math"

const inner_radius = 1
const middle_radius = 5
const outer_radius = 10

const inner_score = 10
const middle_score = 5
const outer_score = 1

// calculatDistanceToOrigin calculates the distance between the origin and the given point
func calculatDistanceToOrigin(x float64, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// Score returns the score of a given shot. The shot is represented as the x, y
// position where the dart landed
func Score(x float64, y float64) int {
	radius := calculatDistanceToOrigin(x, y)

	if radius <= inner_radius {
		return inner_score
	}
	if radius <= middle_radius {
		return middle_score
	}
	if radius <= outer_radius {
		return outer_score
	}

	return 0
}
