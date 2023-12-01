package main

import tea "github.com/charmbracelet/bubbletea"

func main() {
	m := NewModel()
	program := tea.NewProgram(m)
	_, err := program.Run()
	if err != nil {
		panic(err)
	}
}

type Model struct {
	title string
}

func NewModel() Model {
	return Model{
		title: "Hello, World!",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.title
}
