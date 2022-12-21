package voronoi

import (
	"image/color"
	"math/rand"
)

func getRandomColor() color.RGBA {
	return color.RGBA{
		R: getRandomUint8(),
		G: getRandomUint8(),
		B: getRandomUint8(),
		A: 255,
	}
}

func getGrayScaleColor(i, of int) color.RGBA {
	col := uint8(float64(i) / float64(of) * 256)
	return color.RGBA{
		R: col,
		G: col,
		B: col,
		A: 255,
	}
}

func getRandomUint8() uint8 {
	return uint8(rand.Intn(256))
}
