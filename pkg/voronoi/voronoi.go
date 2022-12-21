package voronoi

import (
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Voronoi struct {
	Config
	seeds []seedPoint
	image *image.RGBA
}

func NewVoronoi(config Config) *Voronoi {
	v := &Voronoi{Config: config}
	return v
}

func (v *Voronoi) Generate() {
	v.generateSeedPoints()
	v.generateVoronoi()
	if v.RenderSeedPoints {
		v.drawSeedPoints()
	}
}

func (v *Voronoi) SaveToPng(path string, format ImageFormat) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	switch format {
	case ImageFormatPNG:
		err = png.Encode(file, v.image)
	case ImageFormatGIF:
		err = gif.Encode(file, v.image, nil)
	}
	if err != nil {
		return err
	}

	return nil
}

//
// Private functions
//

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
	// Create image for Voronoi
	v.image = image.NewRGBA(image.Rect(0, 0, v.Width, v.Height))

	// Generate Voronoi
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
