# VORONOI

Package for generating Voronoi diagrams.

## Documentation
**Configuration options**:
* Width and height
* Distance method : Euclidean or Manhattan
* Number of seed points
* Seed point radius and color
* Render seed point or not
* Color scheme : Random or grey scale

## Usage
```Go
func main() {
    v := voronoi.NewVoronoi(
        voronoi.WithSize(800, 600),
        voronoi.WithSeed(100, 5, BLACK, true),
        voronoi.WithScheme(voronoi.ColorSchemeRandom),
        voronoi.WithMethod(voronoi.DistanceMethodManhattan),
    )
	v := voronoi.NewVoronoi(c)
	v.Generate()
	err := v.SaveToPng("/home/per/temp/voronoi.png", voronoi.ImageFormatPNG)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to save to file : %v", err)
	}
}
```
## Links
* Inspiration from Tsoding Daily : https://www.youtube.com/watch?v=kT-Mz87-HcQ
* Wikipedia : https://en.wikipedia.org/wiki/Voronoi_diagram

## Todo

## Screenshots

Euclidean distance:
![screenshot](assets/screenshot.png)
Manhattan distance:
![screenshot](assets/screenshot_manhattan.png)
