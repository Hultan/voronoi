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

func getRandomUint8() uint8 {
	return uint8(rand.Intn(256))
}
