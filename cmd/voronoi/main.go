package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hultan/voronoi/pkg/voronoi"
)

var BLACK = color.RGBA{A: 255}

func main() {
	c := voronoi.NewConfig()
	c.Width = 800
	c.Height = 600
	c.DistanceMethod = voronoi.DistanceMethodEuclidean
	c.NumSeedPoints = 100
	c.SeedPointRadius = 5
	c.SeedPointColor = BLACK
	c.ColorScheme = voronoi.ColorSchemeRandom
	c.RenderSeedPoints = true
	v := voronoi.NewVoronoi(c)
	v.Generate()
	err := v.SaveToPng("/home/per/temp/voronoi.png", voronoi.ImageFormatPNG)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to save to file : %v", err)
	}
}
