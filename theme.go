package main

import (
	_ "embed"
	"image/color"
	"log"

	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"gioui.org/widget/material"
)

//go:embed assets/BigBlueTerm437NerdFontMono-Regular.ttf
var fontTTF []byte

const appFace = "BigBlueTerm"

var palette = struct {
	bg0, bg1     color.NRGBA
	fg1, gray    color.NRGBA
	red, green   color.NRGBA
	yellow, blue color.NRGBA
	aqua, purple color.NRGBA
}{
	bg0:    hexColor(0x282828),
	bg1:    hexColor(0x3c3836),
	fg1:    hexColor(0xebdbb2),
	gray:   hexColor(0xa89984),
	red:    hexColor(0xfb4934),
	green:  hexColor(0xb8bb26),
	yellow: hexColor(0xfabd2f),
	blue:   hexColor(0x83a598),
	aqua:   hexColor(0x8ec07c),
	purple: hexColor(0xd3869b),
}

func newTheme() *material.Theme {
	face, err := opentype.Parse(fontTTF)
	if err != nil {
		log.Fatalf("parse font: %v", err)
	}

	collection := append(gofont.Collection(), font.FontFace{
		Font: font.Font{Typeface: appFace},
		Face: face,
	})

	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(collection))
	th.Face = appFace
	th.Palette.Bg = palette.bg0
	th.Palette.Fg = palette.fg1
	th.Palette.ContrastBg = palette.yellow
	th.Palette.ContrastFg = palette.bg0
	return th
}
