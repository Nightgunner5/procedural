// Package noise implements a perlin noise generator.
package noise

import (
	"math"
	"math/rand"
)

type (
	octave [256]uint8
	Noise  []octave
)

func New(seed int64, octaves int) Noise {
	n := make(Noise, octaves)
	r := rand.New(rand.NewSource(seed))

	for i := range n {
		perm := r.Perm(len(n[i]))

		for j := range n[i] {
			n[i][j] = uint8(perm[j])
		}
	}

	return n
}

func (n Noise) Noise(x, y, z float64) float64 {
	mul := 1.0
	cur := 0.0
	max := 0.0

	for i := range n {
		cur += n[i].noise(x/mul, y/mul, z/mul) * mul
		max += mul
		mul *= 0.25
	}

	return cur / max
}

func floor(f float64) (mod uint8, frac float64) {
	fl := math.Floor(f)
	return uint8(fl), f - fl
}

// Adapted from http://mrl.nyu.edu/~perlin/noise/
func (n octave) noise(x, y, z float64) float64 {
	// Find unit cube that contains point.
	// Find relative X, Y, Z of point in cube.
	X, x := floor(x)
	Y, y := floor(y)
	Z, z := floor(z)

	// Compute fade curves for each of X, Y, Z.
	u, v, w := fade(x), fade(y), fade(z)

	// Hash coordinates of the 8 cube corners,
	var (
		A  = n[X] + Y
		AA = n[A] + Z
		AB = n[A+1] + Z
		B  = n[X+1] + Y
		BA = n[B] + Z
		BB = n[B+1] + Z
	)

	// and add blended results from 8 corners of cube.
	return lerp(w, lerp(v, lerp(u,
		grad(n[AA], x, y, z),
		grad(n[BA], x-1, y, z)),
		lerp(u,
			grad(n[AB], x, y-1, z),
			grad(n[BB], x-1, y-1, z))),
		lerp(v, lerp(u,
			grad(n[AA+1], x, y, z-1),
			grad(n[BA+1], x-1, y, z-1)),
			lerp(u,
				grad(n[AB+1], x, y-1, z-1),
				grad(n[BB+1], x-1, y-1, z-1))))
}
func fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}
func lerp(t, a, b float64) float64 {
	return a + t*(b-a)
}
func grad(hash uint8, x, y, z float64) float64 {
	// Convert low four bits of hash code into twelve gradient directions.
	switch uint(hash & 15) {
	case 0:
		return x + y
	case 1:
		return -x + y
	case 2:
		return x - y
	case 3:
		return -x - y
	case 4:
		return x + z
	case 5:
		return -x + z
	case 6:
		return x - z
	case 7:
		return -x - z
	case 8:
		return y + z
	case 9:
		return -y + z
	case 10:
		return y - z
	case 11:
		return -y - z
	case 12:
		return y + x
	case 13:
		return -y + z
	case 14:
		return y - x
	case 15:
		return -y - z
	}
	panic("unreachable")
}
