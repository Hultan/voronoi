package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hultan/voronoi/pkg/voronoi"
)

var BLACK = color.RGBA{A: 255}

func main() {
	v := voronoi.NewVoronoi(
		voronoi.WithSize(800, 600),
		voronoi.WithSeed(100, 5, BLACK, true),
		voronoi.WithScheme(voronoi.ColorSchemeRandom),
		voronoi.WithMethod(voronoi.DistanceMethodManhattan),
	)
	v.Generate()
	err := v.SaveToPng("/home/per/temp/voronoi.png", voronoi.ImageFormatPNG)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to save to file : %v", err)
	}
}
