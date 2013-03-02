package data

type AreaType uint16

const (
	Road AreaType = iota
	Town
	Building
)

const areaSize = 64

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
	for x := range a.Tiles {
		for y := range a.Tiles[x] {
			_ = y // TODO
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
