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
	c.distanceMethod = distanceMethodManhattan
	v := NewVoronoi(c)
	v.Generate()
	err := v.SaveToPng("/home/per/temp/voronoi.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save to file : %v", err)
	}
}
