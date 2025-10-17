package model

import (
	"fmt"

	"github.com/Yahar4/internal/quotes"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	primaryColor   = lipgloss.Color("#FFD700")
	secondaryColor = lipgloss.Color("#8B4513")
	accentColor    = lipgloss.Color("#FF6B6B")
	textColor      = lipgloss.Color("#F5F5F5")

	titleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			MarginBottom(1)

	quoteStyle = lipgloss.NewStyle().
			Foreground(textColor).
			Italic(true).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(secondaryColor).
			Align(lipgloss.Center).
			Width(60)

	authorStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			MarginTop(1)

	instructionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#AAAAAA")).
				Faint(true).
				MarginTop(1)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF4444")).
			Bold(true).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF4444")).
			Align(lipgloss.Center)

	loadingStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true)
)

type Model struct {
	quotes   []string
	current  string
	loading  bool
	err      error
	filePath string
	width    int
	height   int
}

type quotesLoadedMsg struct {
	quotes []string
	err    error
}

func InitialModel(filePath string) Model {
	return Model{
		loading:  true,
		filePath: filePath,
	}
}

func (m Model) Init() tea.Cmd {
	return loadQuotes(m.filePath)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "r", " ", "enter", "n":
			if len(m.quotes) > 0 {
				m.current = quotes.GetRandomQuote(m.quotes)
			}
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case quotesLoadedMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.quotes = msg.quotes
			if len(msg.quotes) > 0 {
				m.current = quotes.GetRandomQuote(msg.quotes)
			}
		}
		return m, nil
	}

	return m, nil
}

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}

	content := m.renderContent()

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

func (m Model) renderContent() string {
	if m.loading {
		return loadingStyle.Render("🎸 Loading Kanye wisdom...")
	}

	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("❌ Error: %v", m.err))
	}

	if m.current == "" {
		return errorStyle.Render("😕 No quotes available")
	}

	quoteBox := quoteStyle.Render(m.current)
	instructions := instructionStyle.Render("Press r, space, or enter for new quote • Press q to quit")

	return lipgloss.JoinVertical(
		lipgloss.Center,
		titleStyle.Render("🎤 KANYE WEST"),
		quoteBox,
		authorStyle.Render("— Ye"),
		instructions,
	)
}

func loadQuotes(filePath string) tea.Cmd {
	return func() tea.Msg {
		quotes, err := quotes.GetQuotes(filePath)
		return quotesLoadedMsg{quotes: quotes, err: err}
	}
}
