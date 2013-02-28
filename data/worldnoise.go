package data

import (
	"github.com/Nightgunner5/procedural/noise"
)

type Noise struct {
	Elevation  noise.Noise
	Vegetation noise.Noise
	AlignmentV noise.Noise
	AlignmentH noise.Noise
}

func (n *Noise) Init(seed int64) {
	const (
		ElevationOctaves  = 3
		VegetationOctaves = 3
		AlignmentVOctaves = 2
		AlignmentHOctaves = 2
	)

	octaves := noise.New(seed, ElevationOctaves+VegetationOctaves+AlignmentVOctaves+AlignmentHOctaves)

	n.Elevation, octaves = octaves[:ElevationOctaves], octaves[ElevationOctaves:]
	n.Vegetation, octaves = octaves[:VegetationOctaves], octaves[VegetationOctaves:]
	n.AlignmentV, octaves = octaves[:AlignmentVOctaves], octaves[AlignmentVOctaves:]
	n.AlignmentH, octaves = octaves[:AlignmentHOctaves], octaves[AlignmentHOctaves:]
}
