package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Modelo customizado para encapsular o spinner do pacote 'bubbles/spinner'
type Model struct {
	Spinner spinner.Model
	done    bool
}

// Função para criar um novo spinner
func New() Model {
	return Model{
		Spinner: spinner.New(spinner.WithSpinner(spinner.Dot)), // Usando WithSpinner como opção
	}
}

// Função de inicialização
func (m Model) Init() tea.Cmd {
	return tea.Batch(m.Spinner.Tick)
}

// Função para atualizar o spinner
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

// Função de visualização do spinner
func (m Model) View() string {
	if m.done {
		return "Compilation complete!"
	}
	return m.Spinner.View() + " Compiling..."
}

// Função para definir o status de 'done'
func (m *Model) SetDone(done bool) {
	m.done = done
}

// Função para iniciar o spinner
func (m *Model) Start() tea.Cmd {
	return m.Spinner.Tick
}

// Função para alterar o tipo de spinner
func (m *Model) SetSpinnerType(s spinner.Spinner) {
	m.Spinner.Spinner = s
}
