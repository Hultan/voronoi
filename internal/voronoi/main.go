package voronoi

import (
	"fmt"
	"image/color"
	"os"
)

var BLACK = color.RGBA{A: 255}

// GenerateVoronoi : Startup function
func GenerateVoronoi() {
	c := NewConfig()
	c.width = 4000
	c.height = 3000
	c.distanceMethod = distanceMethodEuclidean
	c.numSeedPoints = 100
	c.seedPointRadius = 20
	c.seedPointColor = BLACK
	c.colorScheme = colorSchemeRandom
	v := NewVoronoi(c)
	v.Generate()
	err := v.SaveToPng("/home/per/temp/voronoi.png")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to save to file : %v", err)
	}
}
