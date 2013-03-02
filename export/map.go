package export

import (
	"github.com/Nightgunner5/procedural/data"
	"image"
	"image/color"
	"image/draw"
)

func genMap(a *data.Area) *image.RGBA {
	r := data.NewRand(int64(a.ID))
	img := image.NewRGBA(image.Rect(0, 0, len(a.Tiles)<<2, len(a.Tiles[0])<<2))

	for x := range a.Tiles {
		for y := range a.Tiles[x] {
			t := a.Tiles[x][y]
			if t == nil {
				continue
			}

			draw.Draw(img, image.Rect(x<<2, y<<2, (x+1)<<2, (y+1)<<2), image.NewUniform(color.RGBA{t.Terrain.R, t.Terrain.G, t.Terrain.B, 255}), image.ZP, draw.Src)
			for _, o := range t.Objects {
				for i := 0; i < 16; i++ {
					img.SetRGBA(x<<2+r.Intn(1<<2), y<<2+r.Intn(1<<2), color.RGBA{o.R, o.G, o.B, 255})
				}
			}
		}
	}

	return img
}
