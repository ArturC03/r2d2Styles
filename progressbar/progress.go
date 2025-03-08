package progressbar

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	progress int
	quit     chan struct{}
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	select {
	case <-m.quit:
		return m, tea.Quit
	default:
	}

	switch msg.(type) {
	case tickMsg:
		if m.progress >= 100 {
			return m, tea.Quit
		}
		m.progress += 2
		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	bar := "["
	for i := 0; i < 50; i++ {
		if i < m.progress/2 {
			bar += "="
		} else {
			bar += " "
		}
	}
	bar += fmt.Sprintf("] %d%%", m.progress)
	return bar
}

type tickMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

var (
	p    *tea.Program
	quit chan struct{}
)

// Start inicia a barra de progresso
func Start() {
	quit = make(chan struct{})
	p = tea.NewProgram(model{quit: quit})
	go func() {
		if _, err := p.Run(); err != nil {
			fmt.Println("Erro ao iniciar a barra de progresso:", err)
		}
	}()
}

// Stop finaliza a barra de progresso
func Stop() {
	if quit != nil {
		close(quit)
	}
	if p != nil {
		p.Quit()
	}
}
