package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nightails/fireflight/internal/elgato"
)

// getDevice fetches the device metadata from the Elgato Key Light
func getDevice(url string) tea.Cmd {
	return func() tea.Msg {
		d, err := elgato.GetDevice(url)
		if err != nil {
			return errMsg(err)
		}
		return deviceMsg{Device: d}
	}
}
