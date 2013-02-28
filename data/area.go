package data

type Area struct {
	ID uint64

	Version uint64
	Seed    int64

	Rand *Rand
}
