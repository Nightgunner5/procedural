package export

import (
	"github.com/Nightgunner5/procedural/data"
	"image"
	"image/color"
	"image/draw"
)

func genMap(a *data.Area) *image.RGBA {
	r := data.NewRand(int64(a.ID))
	img := image.NewRGBA(image.Rect(0, -1<<2, len(a.Tiles)<<3, len(a.Tiles[0])<<2))

	for x := range a.Tiles {
		for y := range a.Tiles[x] {
			t := a.Tiles[x][y]
			if t == nil || t.Terrain == nil {
				continue
			}

			if t.Teleport != nil {
				draw.Draw(img, image.Rect(x<<3, (y-1)<<2, (x+1)<<3, (y+1)<<2), image.Black, image.ZP, draw.Src)
				draw.Draw(img, image.Rect(x<<3+1, (y-1)<<2+1, (x+1)<<3-1, (y+1)<<2-1), image.White, image.ZP, draw.Src)
				continue
			}

			draw.Draw(img, image.Rect(x<<3, y<<2, (x+1)<<3, (y+1)<<2), image.NewUniform(color.RGBA{t.Terrain.R, t.Terrain.G, t.Terrain.B, 255}), image.ZP, draw.Src)
			if !t.Terrain.Passable {
				draw.Draw(img, image.Rect(x<<3, (y-1)<<2, (x+1)<<3, y<<2), image.NewUniform(color.RGBA{
					uint8(uint(t.Terrain.R) * 7 / 8),
					uint8(uint(t.Terrain.G) * 7 / 8),
					uint8(uint(t.Terrain.B) * 7 / 8),
					255,
				}), image.ZP, draw.Src)
			}
			for _, o := range t.Objects {
				for i := 0; i < 8; i++ {
					_x := x<<3 + r.Intn(1<<3)
					_y := y<<2 + r.Intn(1<<2)
					img.SetRGBA(_x, _y, color.RGBA{o.R, o.G, o.B, 255})
					img.SetRGBA(_x, _y-1, color.RGBA{o.R, o.G, o.B, 255})
				}
			}
		}
	}

	return img
}
