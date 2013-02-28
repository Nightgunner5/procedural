package data

import (
	"encoding/gob"
	"github.com/Nightgunner5/procedural/noise"
	"math/rand"
)

func init() {
	gob.Register(rand.NewSource(0))
}

func NewRand(seed int64) *Rand {
	return &Rand{rand.NewSource(seed)}
}

type Rand struct {
	Source rand.Source
}

// Rand returns a new *Rand with a seed given by Int63().
func (r *Rand) Rand() *Rand {
	return NewRand(r.Int63())
}

// Noise returns a perlin noise generator with the given number of octaves.
func (r *Rand) Noise(octaves int) noise.Noise {
	n := make(noise.Noise, octaves)

	for i := range n {
		perm := r.Perm(len(n[i]))

		for j := range n[i] {
			n[i][j] = uint8(perm[j])
		}
	}

	return n
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand) Int63() int64 { return r.Source.Int63() }

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32 { return uint32(r.Int63() >> 31) }

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r *Rand) Int31() int32 { return int32(r.Int63() >> 32) }

// Int returns a non-negative pseudo-random int.
func (r *Rand) Int() int {
	u := uint(r.Int63())
	return int(u << 1 >> 1) // clear sign bit if int == int32
}

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	max := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := r.Int63()
	for v > max {
		v = r.Int63()
	}
	return v % n
}

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	max := int32((1 << 31) - 1 - (1<<31)%uint32(n))
	v := r.Int31()
	for v > max {
		v = r.Int31()
	}
	return v % n
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if n <= 1<<31-1 {
		return int(r.Int31n(int32(n)))
	}
	return int(r.Int63n(int64(n)))
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float64() float64 { return float64(r.Int63()) / (1 << 63) }

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float32() float32 { return float32(r.Float64()) }

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func (r *Rand) Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		m[i] = i
	}
	for i := 0; i < n; i++ {
		j := r.Intn(i + 1)
		m[i], m[j] = m[j], m[i]
	}
	return m
}
