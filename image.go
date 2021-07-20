package heatm

import (
	"image"
	"image/color"
)

var _ image.Image = (*HeatMapImage)(nil)

type HeatMapImage struct {
	heatmap HeatMap
	max     int
}

func (h *HeatMapImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (h *HeatMapImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: h.heatmap.Columns,
			Y: h.heatmap.Rows,
		},
	}
}

func (h *HeatMapImage) At(x, y int) color.Color {
	if len(h.heatmap.CountsInts2D[y]) < h.heatmap.Columns {
		return color.RGBA{
			A: 255,
		}
	}
	cnt := int(h.heatmap.CountsInts2D[y][x]) & (1<<24 - 1)
	r := cnt >> 16
	g := cnt >> 8
	b := cnt

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}
}
