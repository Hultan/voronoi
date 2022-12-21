package voronoi

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"sync"
	"time"
)

type colorScheme int

const (
	colorSchemeRandom colorScheme = iota
	colorSchemeGreyScale
)

type distanceMethod int

const (
	distanceMethodEuclidean distanceMethod = iota
	distanceMethodManhattan
)

type Config struct {
	SeedPointConfig
	width, height  int
	colorScheme    colorScheme
	distanceMethod distanceMethod
}

type SeedPointConfig struct {
	numSeedPoints    int
	renderSeedPoints bool
	seedPointColor   color.RGBA
	seedPointRadius  int
}

type Voronoi struct {
	Config
	seeds []SeedPoint
	image *image.RGBA
}

func NewConfig() Config {
	return Config{
		SeedPointConfig: SeedPointConfig{
			numSeedPoints:    30,
			renderSeedPoints: true,
			seedPointColor:   color.RGBA{A: 255},
			seedPointRadius:  5,
		},
		colorScheme:    colorSchemeRandom,
		distanceMethod: distanceMethodEuclidean,
	}
}

func NewVoronoi(config Config) *Voronoi {
	v := &Voronoi{Config: config}
	v.generateSeedPoints()
	return v
}

func (v *Voronoi) Generate() {
	v.generateVoronoi()
	if v.renderSeedPoints {
		v.drawSeedPoints()
	}
}

func (v *Voronoi) SaveToPng(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	err = png.Encode(file, v.image)
	if err != nil {
		return err
	}
	return nil
}

func (v *Voronoi) generateSeedPoints() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < v.numSeedPoints; i++ {
		x, y := rand.Intn(v.width), rand.Intn(v.height)

		// Get color
		var col color.RGBA
		switch v.colorScheme {
		case colorSchemeRandom:
			col = getRandomColor()
		case colorSchemeGreyScale:
			col = getGrayScaleColor(i, v.numSeedPoints)
		}

		v.seeds = append(
			v.seeds, SeedPoint{
				Point: Point{x, y},
				color: col,
			},
		)
	}
}

func (v *Voronoi) generateVoronoi() {
	v.image = image.NewRGBA(image.Rect(0, 0, v.width, v.height))
	wg := sync.WaitGroup{}
	wg.Add(v.width)
	for x := 0; x < v.width; x++ {
		go func(x int) {
			for y := 0; y < v.height; y++ {
				p := Point{x, y}
				closest := p.getClosestSeedPoint(v.distanceMethod, v.seeds)
				v.image.SetRGBA(p.x, p.y, v.seeds[closest].color)
			}
			wg.Done()
		}(x)
	}
	wg.Wait()
}

func (v *Voronoi) drawSeedPoints() {
	for _, seed := range v.seeds {
		v.drawCircle(seed.Point)
	}
}

func (v *Voronoi) drawCircle(c Point) {
	r := v.seedPointRadius
	x0, y0 := c.x-r, c.y-r // Top left corner
	x1, y1 := c.x+r, c.y+r // Bottom right corner

	for x := x0; x < x1; x++ {
		for y := y0; y < y1; y++ {
			dx, dy := x-c.x, y-c.y
			if dx*dx+dy*dy <= r*r {
				v.image.SetRGBA(x, y, v.seedPointColor)
			}
		}
	}
}
