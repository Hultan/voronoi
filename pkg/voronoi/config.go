package voronoi

import (
	"image/color"
)

type ConfigFunc func(*Config)

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

func defaultConfig() Config {
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

func WithSize(width, height int) ConfigFunc {
	return func(cfg *Config) {
		cfg.Width = width
		cfg.Height = height
	}
}

func WithSeed(num, radius int, col color.RGBA, render bool) ConfigFunc {
	return func(cfg *Config) {
		cfg.NumSeedPoints = num
		cfg.SeedPointRadius = radius
		cfg.SeedPointColor = col
		cfg.RenderSeedPoints = render
	}
}

func WithScheme(scheme ColorScheme) ConfigFunc {
	return func(cfg *Config) {
		cfg.ColorScheme = scheme
	}
}

func WithMethod(method DistanceMethod) ConfigFunc {
	return func(cfg *Config) {
		cfg.DistanceMethod = method
	}
}
