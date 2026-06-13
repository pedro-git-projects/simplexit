package main

import (
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// systemctlCmd returns a command factory for "systemctl <verb>"
func systemctlCmd(verb string) func() *exec.Cmd {
	return func() *exec.Cmd {
		return exec.Command("systemctl", verb)
	}
}

// logoutCmd detects the running WM/DE and returns the appropriate kill command.
func logoutCmd() *exec.Cmd {
	desktop := strings.ToLower(strings.TrimSpace(os.Getenv("DESKTOP_SESSION")))
	if idx := strings.LastIndex(desktop, "/"); idx >= 0 {
		desktop = desktop[idx+1:]
	}
	if shell := wmShutdownCmd(desktop); shell != "" {
		return exec.Command("sh", "-c", shell)
	}
	// fallback for unrecognised sessions.
	if id := os.Getenv("XDG_SESSION_ID"); id != "" {
		return exec.Command("loginctl", "terminate-session", id)
	}
	return exec.Command("loginctl", "terminate-user", currentUser())
}

// wmShutdownCmd maps a desktop/WM name to the shell command that ends it
func wmShutdownCmd(desktop string) string {
	switch desktop {
	case "dwm":
		return "pkill dwm"
	case "bspwm":
		return "pkill bspwm"
	case "openbox":
		return "pkill openbox"
	case "i3", "i3-with-shmlog":
		return "pkill i3"
	case "awesome":
		return "pkill awesome"
	case "xmonad":
		return "pkill xmonad"
	case "qtile":
		return "pkill qtile"
	case "spectrwm":
		return "pkill spectrwm"
	case "herbstluftwm":
		return "herbstclient quit"
	case "jwm":
		return "pkill jwm"
	case "leftwm":
		return "pkill leftwm"
	case "cwm":
		return "pkill cwm"
	case "fvwm3":
		return "pkill fvwm3"
	case "icewm", "icewm-session":
		return "pkill icewm"
	case "berry":
		return "pkill berry"
	case "worm":
		return "pkill worm"
	case "dk":
		return "dkcmd exit"
	case "dusk":
		return "pkill dusk"
	case "nimdow":
		return "pkill nimdow"
	case "gnome", "gnome-xorg", "gnome-classic":
		return "gnome-session-quit --logout --no-prompt"
	case "xfce":
		return "xfce4-session-logout -f -l"
	case "lxqt":
		return "pkill lxqt"
	case "sway":
		return "pkill sway"
	case "hyprland":
		return "hyprctl dispatch exit"
	case "river":
		return "pkill river"
	case "wayfire":
		return "pkill wayfire"
	default:
		return ""
	}
}

// lockCmd picks the first available screen locker from a priority list
func lockCmd() *exec.Cmd {
	lockers := []struct {
		name string
		args []string
	}{
		{"slock", nil},
		{"i3lock", nil},
		{"betterlockscreen", []string{"-l"}},
	}
	for _, l := range lockers {
		if _, err := exec.LookPath(l.name); err == nil {
			return exec.Command(l.name, l.args...)
		}
	}
	return exec.Command("loginctl", "lock-session")
}

func currentUser() string {
	if u, err := user.Current(); err == nil {
		return u.Username
	}
	return os.Getenv("USER")
}
