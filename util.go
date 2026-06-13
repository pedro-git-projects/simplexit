package main

import (
	"image"
	"image/color"

	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func fillRect(ops *op.Ops, size image.Point, col color.NRGBA) {
	defer clip.Rect{Max: size}.Push(ops).Pop()
	paint.ColorOp{Color: col}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

// hexColor converts a 24-bit RGB integer to an opaque NRGBA value
func hexColor(rgb uint32) color.NRGBA {
	return color.NRGBA{
		R: uint8(rgb >> 16),
		G: uint8(rgb >> 8),
		B: uint8(rgb),
		A: 0xff,
	}
}
