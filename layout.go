package main

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

const (
	outerPadding  = unit.Dp(28)
	gridGap       = unit.Dp(12)
	buttonPadding = unit.Dp(12)
	buttonRadius  = unit.Dp(8)
	gridColumns   = 3
)

func layoutRoot(gtx layout.Context, th *material.Theme, actions []action, status string) layout.Dimensions {
	return layout.UniformInset(outerPadding).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Middle,
			Spacing:   layout.SpaceSides,
		}.Layout(gtx,
			layout.Rigid(centeredLabel(material.H4(th, "init 0 interface"), palette.yellow)),
			layout.Rigid(vSpacer(8)),
			layout.Rigid(centeredLabel(material.Body1(th, currentUser()), palette.gray)),
			layout.Rigid(vSpacer(26)),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return actionGrid(gtx, th, actions)
			}),
			layout.Rigid(vSpacer(24)),
			layout.Rigid(centeredLabel(material.Body2(th, status), palette.gray)),
		)
	})
}

// centeredLabel returns a widget that renders lbl with the given color,
// centered horizontally
func centeredLabel(lbl material.LabelStyle, col color.NRGBA) layout.Widget {
	lbl.Color = col
	lbl.Alignment = text.Middle
	return lbl.Layout
}

// actionGrid arranges the action buttons in rows of gridColumns
func actionGrid(gtx layout.Context, th *material.Theme, actions []action) layout.Dimensions {
	var rows []layout.FlexChild
	for i := 0; i < len(actions); i += gridColumns {
		if i > 0 {
			rows = append(rows, layout.Rigid(vSpacer(gridGap)))
		}
		start, end := i, min(i+gridColumns, len(actions))
		rows = append(rows, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return actionRow(gtx, th, actions[start:end])
		}))
	}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, rows...)
}

// actionRow lays out a horizontal slice of action buttons with uniform spacing
func actionRow(gtx layout.Context, th *material.Theme, row []action) layout.Dimensions {
	children := make([]layout.FlexChild, 0, len(row)*2-1)
	for i := range row {
		if i > 0 {
			children = append(children, layout.Rigid(hSpacer(gridGap)))
		}
		a := &row[i]
		children = append(children, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return actionButton(gtx, th, a)
		}))
	}
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, children...)
}

func vSpacer(h unit.Dp) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Spacer{Height: h}.Layout(gtx)
	}
}

func hSpacer(w unit.Dp) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Spacer{Width: w}.Layout(gtx)
	}
}
