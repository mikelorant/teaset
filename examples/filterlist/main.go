package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/teaset/filterlist"
)

type Model struct {
	filterlist filterlist.Model
}

type Item struct {
	title string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return "" }
func (i Item) FilterValue() string { return "" }

var items = []Item{
	{title: "Red"},
	{title: "Green"},
	{title: "Yellow"},
	{title: "Blue"},
	{title: "Magenta"},
	{title: "Cyan"},
}

var (
	red     = lipgloss.Color("1")
	green   = lipgloss.Color("2")
	yellow  = lipgloss.Color("3")
	blue    = lipgloss.Color("4")
	magenta = lipgloss.Color("5")
	cyan    = lipgloss.Color("6")
)

const (
	defaultWidth         = 40
	defaultHeight        = 5
	defaultItemIndicator = "#"
	defaultPlaceholder   = "Placeholder"
	defaultPromptMark    = "*"
	defaultPromptText    = "Choose colour:"
)

func main() {
	p := tea.NewProgram(New())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

func New() Model {
	fl := filterlist.New()
	fl = Settings(fl)
	fl = Styles(fl)
	fl.SetItems(filterlist.ToItems(items))
	fl.Focus()

	return Model{
		filterlist: fl,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		//nolint:gocritic
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "enter":
			return m, tea.Quit
		}
	}

	m.filterlist, cmd = m.filterlist.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.filterlist.View()
}

func Settings(fl filterlist.Model) filterlist.Model {
	fl.Width = defaultWidth
	fl.Height = defaultHeight
	fl.List.Styles.ItemIndicator = defaultItemIndicator
	fl.TextInput.Placeholder = defaultPlaceholder
	fl.TextInput.PromptMark = defaultPromptMark
	fl.TextInput.PromptText = defaultPromptText

	return fl
}

func Styles(fl filterlist.Model) filterlist.Model {
	fl.List.Styles.Item = lipgloss.NewStyle().Foreground(red)
	fl.List.Styles.ItemSelected = lipgloss.NewStyle().Foreground(green)
	fl.Paginator.Styles.DotEmpty = lipgloss.NewStyle().Foreground(magenta)
	fl.Paginator.Styles.DotFilled = lipgloss.NewStyle().Foreground(cyan)
	fl.TextInput.Styles.PromptText = lipgloss.NewStyle().Foreground(yellow)
	fl.TextInput.Styles.Text = lipgloss.NewStyle().Foreground(blue)

	return fl
}
