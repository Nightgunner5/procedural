package data

type World struct {
	Version uint64
	Seed    int64

	Rand *Rand

	AreaCount   uint64
	CurrentArea uint64

	__TODO__areacache map[uint64]*Area // TODO: actual saving/loading
}

func (w *World) Init() {
	w.Rand = NewRand(w.Seed)

	w.__TODO__areacache = make(map[uint64]*Area) // TODO: actual saving/loading

	a := w.NewArea(w.Rand, Town)
	a.Generate()
	w.SaveArea(a)
	w.CurrentArea = a.ID
}

func (w *World) Area(index uint64) *Area {
	return w.__TODO__areacache[index] // TODO: actual saving/loading
}

func (w *World) NewArea(r *Rand, t AreaType) *Area {
	var a Area
	a.Rand = r.Rand()
	a.ID = w.AreaCount
	a.Type = t
	a.generateName()
	w.AreaCount++

	w.SaveArea(&a)

	return &a
}

func (w *World) SaveArea(a *Area) {
	w.__TODO__areacache[a.ID] = a // TODO: actual saving/loading
}
