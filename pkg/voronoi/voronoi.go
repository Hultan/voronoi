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

type Voronoi struct {
	Config
	seeds []seedPoint
	image *image.RGBA
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

func NewVoronoi(config Config) *Voronoi {
	v := &Voronoi{Config: config}
	v.generateSeedPoints()
	return v
}

func (v *Voronoi) Generate() {
	v.generateVoronoi()
	if v.RenderSeedPoints {
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

	for i := 0; i < v.NumSeedPoints; i++ {
		x, y := rand.Intn(v.Width), rand.Intn(v.Height)

		// Get color
		var col color.RGBA
		switch v.ColorScheme {
		case ColorSchemeRandom:
			col = getRandomColor()
		case ColorSchemeGreyScale:
			col = getGrayScaleColor(i, v.NumSeedPoints)
		}

		v.seeds = append(
			v.seeds, seedPoint{
				point: point{x, y},
				color: col,
			},
		)
	}
}

func (v *Voronoi) generateVoronoi() {
	v.image = image.NewRGBA(image.Rect(0, 0, v.Width, v.Height))
	wg := sync.WaitGroup{}
	wg.Add(v.Width)
	for x := 0; x < v.Width; x++ {
		go func(x int) {
			for y := 0; y < v.Height; y++ {
				p := point{x, y}
				closest := p.getClosestSeedPoint(v.DistanceMethod, v.seeds)
				v.image.SetRGBA(p.x, p.y, v.seeds[closest].color)
			}
			wg.Done()
		}(x)
	}
	wg.Wait()
}

func (v *Voronoi) drawSeedPoints() {
	for _, seed := range v.seeds {
		v.drawCircle(seed.point)
	}
}

func (v *Voronoi) drawCircle(c point) {
	r := v.SeedPointRadius
	x0, y0 := c.x-r, c.y-r // Top left corner
	x1, y1 := c.x+r, c.y+r // Bottom right corner

	for x := x0; x < x1; x++ {
		for y := y0; y < y1; y++ {
			dx, dy := x-c.x, y-c.y
			if dx*dx+dy*dy <= r*r {
				v.image.SetRGBA(x, y, v.SeedPointColor)
			}
		}
	}
}
