package data

type Tile struct {
	Teleport *Teleport
	Terrain  *Terrain
	Objects  []*Object
}

type Teleport struct {
	Area uint64
	X, Y uint
}

type Terrain struct {
	R, G, B uint8
	// TODO
}

type Object struct {
	R, G, B uint8
	// TODO
}

func NewTerrain(r *Rand) *Terrain {
	return &Terrain{
		R: uint8(r.Intn(256)),
		G: uint8(r.Intn(128)),
		B: uint8(r.Intn(128)),
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
