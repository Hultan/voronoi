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
	colorRandom colorScheme = iota
	colorGreyScale
)

type SeedPointConfig struct {
	numSeedPoints    int
	renderSeedPoints bool
	seedPointColor   color.RGBA
	seedPointRadius  int
	colorScheme      colorScheme
}

type Voronoi struct {
	SeedPointConfig
	seeds []SeedPoint
	image *image.RGBA
}

func NewConfig(num int, col color.RGBA, radius int, scheme colorScheme) SeedPointConfig {
	return SeedPointConfig{
		numSeedPoints:    num,
		renderSeedPoints: true,
		seedPointColor:   col,
		seedPointRadius:  radius,
		colorScheme:      scheme,
	}
}

func NewConfigNoRender(num int) SeedPointConfig {
	return SeedPointConfig{
		numSeedPoints:    num,
		renderSeedPoints: false,
		seedPointColor:   color.RGBA{},
		seedPointRadius:  0,
	}
}

func NewVoronoi(config SeedPointConfig) *Voronoi {
	v := &Voronoi{SeedPointConfig: config}
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
		x := rand.Intn(800)
		y := rand.Intn(600)

		// Get color
		var color color.RGBA
		switch v.colorScheme {
		case colorRandom:
			color = getRandomColor()
		case colorGreyScale:
			getGrayScaleColor(i, v.numSeedPoints)
		}

		v.seeds = append(
			v.seeds, SeedPoint{
				Point: Point{x, y},
				color: color,
			},
		)
	}
}

func (v *Voronoi) generateVoronoi() {
	v.image = image.NewRGBA(image.Rect(0, 0, 800, 600))
	wg := sync.WaitGroup{}
	wg.Add(800)
	for x := 0; x < 800; x++ {
		go func(x int) {
			for y := 0; y < 600; y++ {
				p := Point{x, y}
				v.image.SetRGBA(p.x, p.y, v.seeds[p.getClosestTo(v.seeds)].color)
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
