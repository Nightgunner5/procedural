package data

type World struct {
	Version uint64
	Seed    int64

	Rand *Rand

	AreaCount   uint64
	CurrentArea uint64
}

func (w *World) Init() {
	w.Rand = NewRand(w.Seed)
}

func (w *World) Area(index uint64) *Area {
	return new(Area) // TODO
}
