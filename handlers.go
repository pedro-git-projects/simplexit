package main

import (
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
)

// handleKeys drains the key-event queue. It returns ("", true) when the
// user presses Escape, signalling the caller to exit. Otherwise it returns
// the last error message (or "") and false
func handleKeys(gtx layout.Context, actions []action, filters []event.Filter) (string, bool) {
	var status string
	for {
		ev, ok := gtx.Event(filters...)
		if !ok {
			return status, false
		}
		ke, ok := ev.(key.Event)
		if !ok || ke.State != key.Press {
			continue
		}
		if ke.Name == key.NameEscape {
			return "", true
		}
		for i := range actions {
			if ke.Name == actions[i].key {
				if msg := runAction(actions[i]); msg != "" {
					status = msg
				}
			}
		}
	}
}

// handleClicks drains click events from every action button
func handleClicks(gtx layout.Context, actions []action) string {
	var status string
	for i := range actions {
		for actions[i].button.Clicked(gtx) {
			if msg := runAction(actions[i]); msg != "" {
				status = msg
			}
		}
	}
	return status
}
