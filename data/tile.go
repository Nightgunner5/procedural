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
	// TODO
}

type Object struct {
	// TODO
}
