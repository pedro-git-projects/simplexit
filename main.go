package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(
			app.Title(appTitle),
			app.Size(appWidth, appHeight),
			app.Decorated(false),
		)
		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	theme := newTheme()
	actions := newActions()
	filters := keyFilters(actions)

	var ops op.Ops
	status := "ESC to abort"

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if msg, quit := handleKeys(gtx, actions, filters); quit {
				return nil
			} else if msg != "" {
				status = msg
			}

			if msg := handleClicks(gtx, actions); msg != "" {
				status = msg
			}

			fillRect(gtx.Ops, gtx.Constraints.Max, palette.bg0)
			layoutRoot(gtx, theme, actions, status)
			e.Frame(gtx.Ops)
		}
	}
}
