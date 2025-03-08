package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	spinner spinner.Model
	done    bool
}

func New() Model {
	return Model{
		spinner: spinner.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {
	if m.done {
		return "Compilation complete!"
	}
	return m.spinner.View() + " Compiling..."
}

func (m *Model) SetDone(done bool) {
	m.done = done
}

func (m *Model) Start() tea.Cmd {
	return m.spinner.Tick
}
