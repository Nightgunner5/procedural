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

func (a *Area) Generate(w *World) {
	if a.Generated {
		return
	}

	switch a.Type {
	case Road:
		a.generateRoad(w)
	case Town:
		a.generateTown(w)
	case Building:
		a.generateBuilding(w)
	}

	a.Generated = true

	w.SaveArea(a)
}

func (a *Area) generateRoad(w *World) {
	for x := range a.Tiles {
		for y := range a.Tiles[x] {
			_ = y // TODO
		}
	}
}

func (a *Area) generateTown(w *World) {
	ground := make([]*Terrain, a.Rand.Intn(10)+1)
	for i := range ground {
		ground[i] = NewTerrain(a.Rand)
	}
	groundCover := make([]*Object, a.Rand.Intn(10)+1)
	for i := range groundCover {
		groundCover[i] = NewPlant(a.Rand, false)
	}

	street := NewBuildingTerrain(a.Rand, true)
	building := make([]*Terrain, a.Rand.Intn(3)+1)
	for i := range building {
		building[i] = NewBuildingTerrain(a.Rand, false)
	}

buildingLoop:
	for i := -2; i < 2; i++ {
		for j := 0; j <= 1; j++ {
			mat := building[a.Rand.Intn(len(building))]

			x0 := areaSize/2 + i*10 + 1
			y0 := areaSize/2 - 10*j

			if i < 0 {
				x0--
			} else {
				x0++
			}

			x1 := x0 + 8
			y1 := y0 + 8
			x2 := x0 + 2
			y2 := y1 - 1
			if j == 0 {
				x2 = x1 - 3
				y2 = y0
			}

			for x := x0; x < x1; x++ {
				for y := y0; y < y2; y++ {
					if a.Tiles[x][y] != nil {
						continue buildingLoop
					}
				}
			}
			if j == 1 {
				for x := x0 - 1; x <= x1; x++ {
					a.Tiles[x][y1] = &Tile{
						Terrain: street,
					}
					a.Tiles[x][y1+1] = &Tile{
						Terrain: street,
					}
				}
			}
			for x := x0; x < x1; x++ {
				for y := y0; y < y1; y++ {
					a.Tiles[x][y] = &Tile{
						Terrain: mat,
					}
				}
				for y := y1 + 2; y < y2; y++ {
					a.Tiles[x][y] = &Tile{
						Terrain: mat,
					}
				}
			}
			a.Tiles[x2][y2].Terrain = street
			subarea := w.NewArea(a.Rand, Building)
			a.Tiles[x2][y2].Teleport = &Teleport{
				Area: subarea.ID,
				X:    areaSize / 2,
				Y:    areaSize - 2,
			}
			subarea.Tiles[areaSize/2][areaSize-1] = &Tile{
				Teleport: &Teleport{
					Area: a.ID,
					X:    x2,
					Y:    y2,
				},
			}
			w.SaveArea(subarea)
		}
	}

	for x := areaSize/2 - 1; x < areaSize/2+1; x++ {
		for y := areaSize / 8; y < areaSize*7/8; y++ {
			if a.Tiles[x][y] != nil {
				continue
			}

			if y == areaSize/8 || y == areaSize*7/8-1 {
				if x == areaSize/2 {
					t := a.Tiles[x-1][y].Teleport
					a.Tiles[x][y] = &Tile{
						Terrain: street,
						Teleport: &Teleport{
							Area: t.Area,
							X:    x,
							Y:    t.Y,
						},
					}
					subarea := w.Area(t.Area)
					subarea.Tiles[x][t.Y] = &Tile{
						Terrain: street,
						Teleport: &Teleport{
							Area: a.ID,
							X:    x,
							Y:    y,
						},
					}
					w.SaveArea(subarea)
					continue
				}

				subarea := w.NewArea(a.Rand, Road)
				if y < areaSize/2 {
					a.Tiles[x][y] = &Tile{
						Terrain: street,
						Teleport: &Teleport{
							Area: subarea.ID,
							X:    x,
							Y:    0,
						},
					}
					subarea.Tiles[x][0] = &Tile{
						Terrain: street,
						Teleport: &Teleport{
							Area: a.ID,
							X:    x,
							Y:    y,
						},
					}
				} else {
					a.Tiles[x][y] = &Tile{
						Terrain: street,
						Teleport: &Teleport{
							Area: subarea.ID,
							X:    x,
							Y:    areaSize - 1,
						},
					}
					subarea.Tiles[x][areaSize-1] = &Tile{
						Terrain: street,
						Teleport: &Teleport{
							Area: a.ID,
							X:    x,
							Y:    y,
						},
					}
				}
				w.SaveArea(subarea)
			} else {
				a.Tiles[x][y] = &Tile{
					Terrain: street,
				}
			}
		}
	}

	var (
		noise            = a.Rand.Noise(3)
		areaSizeNoise    = a.Rand.Float64() * 256
		groundTypeNoise  = a.Rand.Float64() * 256
		groundCoverNoise = a.Rand.Float64() * 256
	)

	for x := range a.Tiles {
		for y := range a.Tiles[x] {
			maxSize := int(areaSize/2 - noise.Noise(float64(x)/8, float64(y)/8, areaSizeNoise)*(areaSize/8) - areaSize/8)
			if (x-areaSize/2)*(x-areaSize/2)+(y-areaSize/2)*(y-areaSize/2) > maxSize*maxSize {
				continue
			}

			var t Tile
			if a.Tiles[x][y] != nil {
				t = *a.Tiles[x][y]
			}
			a.Tiles[x][y] = &t

			if t.Terrain != nil {
				continue
			}

			f := noise.Noise(float64(x)/48, float64(y)/48, groundTypeNoise)/2 + 0.5
			t.Terrain = ground[int(f*float64(len(ground)))%len(ground)]

			f = noise.Noise(float64(x)/16, float64(y)/16, groundCoverNoise)/2 + 0.5
			t.Objects = append(t.Objects, groundCover[int(f*float64(len(groundCover)))%len(groundCover)])
		}
	}
}

func (a *Area) generateBuilding(w *World) {
	for x := range a.Tiles {
		for y := range a.Tiles[x] {
			_ = y // TODO
		}
	}
}
