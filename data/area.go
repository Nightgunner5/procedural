package data

type AreaType uint16

const (
	Road AreaType = iota
	Town
	Building
)

const areaSize = 96

type Area struct {
	ID        uint64
	Name      string
	Generated bool

	Version uint64
	Type    AreaType

	Rand *Rand

	Tiles [areaSize][areaSize]*Tile
}

func (a *Area) generateName() {
	a.Name = "Unnamed Area" // TODO: names
}

func (a *Area) Generate() {
	if a.Generated {
		return
	}

	switch a.Type {
	case Road:
		a.generateRoad()
	case Town:
		a.generateTown()
	case Building:
		a.generateBuilding()
	}

	a.Generated = true
}

func (a *Area) generateRoad() {
	for x := range a.Tiles {
		for y := range a.Tiles[x] {
			_ = y // TODO
		}
	}
}

func (a *Area) generateTown() {
	ground := make([]*Terrain, a.Rand.Intn(10)+1)
	for i := range ground {
		ground[i] = NewTerrain(a.Rand)
	}
	groundCover := make([]*Object, a.Rand.Intn(10)+1)
	for i := range groundCover {
		groundCover[i] = NewPlant(a.Rand, false)
	}

	var (
		noise            = a.Rand.Noise(3)
		areaSizeNoise    = a.Rand.Float64() * 256
		groundTypeNoise  = a.Rand.Float64() * 256
		groundCoverNoise = a.Rand.Float64() * 256
	)

	for x := range a.Tiles {
		for y := range a.Tiles[x] {
			maxSize := int(areaSize/2 - noise.Noise(float64(x)/8, float64(y)/8, areaSizeNoise)*(areaSize/8))
			if (x-areaSize/2)*(x-areaSize/2)+(y-areaSize/2)*(y-areaSize/2) > maxSize*maxSize {
				continue
			}

			var t Tile
			a.Tiles[x][y] = &t

			f := noise.Noise(float64(x)/48, float64(y)/48, groundTypeNoise)/2 + 0.5
			t.Terrain = ground[int(f*float64(len(ground)))%len(ground)]

			f = noise.Noise(float64(x)/16, float64(y)/16, groundCoverNoise)/2 + 0.5
			t.Objects = append(t.Objects, groundCover[int(f*float64(len(groundCover)))%len(groundCover)])
		}
	}
}

func (a *Area) generateBuilding() {
	for x := range a.Tiles {
		for y := range a.Tiles[x] {
			_ = y // TODO
		}
	}
}
