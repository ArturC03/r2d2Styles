package progressbar

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	progress progress.Model
	done     bool
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}

	case tickMsg:
		if m.done {
			return m, tea.Quit
		}
		cmd := m.progress.IncrPercent(0.05)
		if m.progress.Percent() >= 1.0 {
			m.done = true
			return m, tea.Quit
		}
		return m, tea.Batch(cmd, tick())
	}

	return m, nil
}

func (m model) View() string {
	if m.done {
		return "Compilação concluída!\n"
	}
	return fmt.Sprintf("Compilando...\n%s", m.progress.View())
}

type tickMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

// Start inicia a barra de progresso
func Start() {
	p := tea.NewProgram(model{progress: progress.New(progress.WithDefaultGradient())})
	if err := p.Start(); err != nil {
		fmt.Println("Erro ao iniciar a barra de progresso:", err)
	}
}
