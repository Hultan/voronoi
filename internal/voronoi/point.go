package voronoi

import (
	"image/color"
	"math"
)

type point struct {
	x, y int
}

type seedPoint struct {
	point
	color color.RGBA
}

func (p point) distanceTo(method DistanceMethod, s point) float64 {
	result := 0.0

	dx1, dy1 := p.x-s.x, p.y-s.y
	switch method {
	case DistanceMethodEuclidean:
		result = math.Sqrt(float64(dx1*dx1 + dy1*dy1))
	case DistanceMethodManhattan:
		result = float64(abs(dx1) + abs(dy1))
	}

	return result
}

func (p point) getClosestSeedPoint(method DistanceMethod, seeds []seedPoint) int {
	i := 0

	for s := 1; s < len(seeds); s++ {
		// If distance to point i is greater than distance to point s,
		// set i to equal s
		if p.distanceTo(method, seeds[i].point) > p.distanceTo(method, seeds[s].point) {
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
