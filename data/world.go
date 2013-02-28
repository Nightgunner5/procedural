package data

type World struct {
	Version uint64
	Seed    int64

	Rand *Rand

	AreaCount uint64
}

func (w *World) Init() {
	w.Rand = NewRand(w.Seed)
}
