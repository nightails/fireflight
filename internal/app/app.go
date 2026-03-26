package app

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nightails/fireflight/internal/elgato"
)

type Model struct {
	URL    string
	Device elgato.Device
	Err    error
}

func NewModel(url string) Model {
	return Model{
		URL:    url,
		Device: elgato.Device{},
	}
}

func (m Model) Init() tea.Cmd {
	return getDevice(m.URL)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	default:
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		default:
			return m, nil
		case "ctrl+c":
			return m, tea.Quit
		}
	case deviceMsg:
		m.Device = msg.Device
		return m, nil
	case errMsg:
		m.Err = msg
		return m, nil
	}
}

func (m Model) View() string {
	b := strings.Builder{}
	b.WriteString("Welcome to Fireflight!\n\n")

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("\n\nError: %s", m.Err))
		return b.String()
	}

	b.WriteString(fmt.Sprintf("Device: %s", m.Device.ProductName))
	return b.String()
}
