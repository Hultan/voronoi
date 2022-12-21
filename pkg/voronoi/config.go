package voronoi

import (
	"image/color"
)

type ImageFormat int

const (
	ImageFormatPNG ImageFormat = iota
	ImageFormatGIF
)

type ColorScheme int

const (
	ColorSchemeRandom ColorScheme = iota
	ColorSchemeGreyScale
)

type DistanceMethod int

const (
	DistanceMethodEuclidean DistanceMethod = iota
	DistanceMethodManhattan
)

type Config struct {
	SeedPointConfig
	Width, Height  int
	ColorScheme    ColorScheme
	DistanceMethod DistanceMethod
}

type SeedPointConfig struct {
	NumSeedPoints    int
	RenderSeedPoints bool
	SeedPointColor   color.RGBA
	SeedPointRadius  int
}

func NewConfig() Config {
	return Config{
		SeedPointConfig: SeedPointConfig{
			NumSeedPoints:    30,
			RenderSeedPoints: true,
			SeedPointColor:   color.RGBA{A: 255},
			SeedPointRadius:  5,
		},
		Width:          800,
		Height:         600,
		ColorScheme:    ColorSchemeRandom,
		DistanceMethod: DistanceMethodEuclidean,
	}
}
