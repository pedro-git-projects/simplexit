package main

import (
	"fmt"
	"image/color"
	"os/exec"

	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type action struct {
	key     key.Name
	label   string
	hint    string
	color   color.NRGBA
	command func() *exec.Cmd
	button  widget.Clickable
}

func newActions() []action {
	return []action{
		{key: "L", label: "[L] Logout", hint: "terminate current session", color: palette.blue, command: logoutCmd},
		{key: "R", label: "[R] Reboot", hint: "systemctl reboot", color: palette.yellow, command: systemctlCmd("reboot")},
		{key: "S", label: "[S] Shutdown", hint: "systemctl poweroff", color: palette.red, command: systemctlCmd("poweroff")},
		{key: "U", label: "[U] Suspend", hint: "systemctl suspend", color: palette.aqua, command: systemctlCmd("suspend")},
		{key: "H", label: "[H] Hibernate", hint: "systemctl hibernate", color: palette.purple, command: systemctlCmd("hibernate")},
		{key: "K", label: "[K] Lock", hint: "slock / i3lock / loginctl", color: palette.green, command: lockCmd},
	}
}

func runAction(a action) string {
	cmd := a.command()
	if cmd == nil {
		return fmt.Sprintf("%s: no command configured", a.label)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Sprintf("%s: %v", a.label, err)
	}
	return ""
}

// keyFilters derives the set of key.Filters from the registered actions,
// keeping the filter list and action definitions in sync automatically
func keyFilters(actions []action) []event.Filter {
	filters := make([]event.Filter, 1, len(actions)+1)
	filters[0] = key.Filter{Name: key.NameEscape}
	for _, a := range actions {
		filters = append(filters, key.Filter{Name: a.key})
	}
	return filters
}

// actionButton renders a single styled button for the given action
func actionButton(gtx layout.Context, th *material.Theme, a *action) layout.Dimensions {
	btn := material.Button(th, &a.button, a.label)
	btn.Background = palette.bg1
	btn.Color = a.color
	btn.CornerRadius = buttonRadius
	btn.Inset = layout.UniformInset(buttonPadding)
	return btn.Layout(gtx)
}
