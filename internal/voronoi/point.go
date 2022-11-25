package voronoi

import (
	"image/color"
)

type Point struct {
	x, y int
}

type SeedPoint struct {
	Point
	color color.RGBA
}

func (p Point) closestTo(s1, s2 Point) bool {
	dx1, dy1 := p.x-s1.x, p.y-s1.y
	dx2, dy2 := p.x-s2.x, p.y-s2.y
	return dx1*dx1+dy1*dy1 <= dx2*dx2+dy2*dy2
}

func (p Point) getClosestTo(seeds []SeedPoint) int {
	i := 0

	for s := 1; s < len(seeds); s++ {
		if !p.closestTo(seeds[i].Point, seeds[s].Point) {
			i = s
		}
	}

	return i
}
