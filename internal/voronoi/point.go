package voronoi

import (
	"image/color"
	"math"
)

type Point struct {
	x, y int
}

type SeedPoint struct {
	Point
	color color.RGBA
}

func (p Point) distanceTo(method distanceMethod, s Point) float64 {
	result := 0.0

	dx1, dy1 := p.x-s.x, p.y-s.y
	switch method {
	case distanceMethodEuclidean:
		result = math.Sqrt(float64(dx1*dx1 + dy1*dy1))
	case distanceMethodManhattan:
		result = float64(abs(dx1) + abs(dy1))
	}

	return result
}

func (p Point) getClosestSeedPoint(method distanceMethod, seeds []SeedPoint) int {
	i := 0

	for s := 1; s < len(seeds); s++ {
		// If distance to point i is greater than distance to point s,
		// set i to equal s
		if p.distanceTo(method, seeds[i].Point) > p.distanceTo(method, seeds[s].Point) {
			i = s
		}
	}

	return i
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
