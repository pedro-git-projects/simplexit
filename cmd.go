package main

import (
	"os"
	"os/exec"
	"os/user"
)

// systemctlCmd returns a command factory for "systemctl <verb>"
func systemctlCmd(verb string) func() *exec.Cmd {
	return func() *exec.Cmd {
		return exec.Command("systemctl", verb)
	}
}

// logoutCmd builds the best available logout command for the current session
func logoutCmd() *exec.Cmd {
	if id := os.Getenv("XDG_SESSION_ID"); id != "" {
		return exec.Command("loginctl", "terminate-session", id)
	}
	if _, err := exec.LookPath("dwm"); err == nil {
		return exec.Command("pkill", "dwm")
	}
	return exec.Command("loginctl", "terminate-user", os.Getenv("USER"))
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
