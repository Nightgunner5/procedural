package data

type Tile struct {
	Teleport *Teleport
	Terrain  *Terrain
	Objects  []*Object
}

type Teleport struct {
	Area uint64
	X, Y int
}

type Terrain struct {
	R, G, B  uint8
	Passable bool
	// TODO
}

type Object struct {
	R, G, B uint8
	// TODO
}

func NewTerrain(r *Rand) *Terrain {
	return &Terrain{
		R:        uint8(r.Intn(256)),
		G:        uint8(r.Intn(128)),
		B:        uint8(r.Intn(128)),
		Passable: true,
		// TODO
	}
}

func NewPlant(r *Rand, significant bool) *Object {
	return &Object{
		R: uint8(r.Intn(256)),
		G: uint8(r.Intn(256)),
		B: uint8(r.Intn(128)),
		// TODO
	}
}

func NewBuildingTerrain(r *Rand, passable bool) *Terrain {
	switch r.Intn(2) {
	case 0:
		// stone
		return &Terrain{
			R:        uint8(r.Intn(32) + 112),
			G:        uint8(r.Intn(32) + 112),
			B:        uint8(r.Intn(32) + 112),
			Passable: passable,
			// TODO
		}
	case 1:
		// wood
		return &Terrain{
			R:        uint8(r.Intn(32) + 200),
			G:        uint8(r.Intn(32) + 128),
			B:        uint8(r.Intn(32) + 64),
			Passable: passable,
			// TODO
		}
	}
	panic("unreachable")
}
