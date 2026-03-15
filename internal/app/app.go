package app

import tea "github.com/charmbracelet/bubbletea"

type Model struct{}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	// TODO: start connect to an elgato device with provided URL and get it's information
	return nil
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
	}
}

func (m Model) View() string {
	return "Welcome to Fireflight!"
}
