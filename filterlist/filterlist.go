package filterlist

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model is the Bubble Tea model for the Filter List widget
type Model struct {
	// Width of the model.
	Width int

	// Height of the model.
	Height int

	// Text input options.
	TextInput TextInput

	// List options.
	List List

	// Paginator options.
	Paginator Paginator

	focus        bool
	textInput    textinput.Model
	list         list.Model
	selectedItem list.Item
}

const (
	defaultWidth  = 20
	defaultHeight = 5
)

// New creates a new Filter List widget.
func New() Model {
	ti := TextInput{
		PromptMark: defaultPromptMark,
		PromptText: defaultPromptText,
	}

	l := List{
		Styles: ListStyles{
			ItemIndicator: defaultItemIndicator,
		},
	}

	return Model{
		List:      l,
		TextInput: ti,
		Width:     defaultWidth,
		Height:    defaultHeight,

		textInput: NewTextInput(ti),
		list:      NewList(l),
	}
}

// Updaqte is the Bubble Tea update loop.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	cmds = append(cmds, m.setModels())

	if !m.focus {
		return m, tea.Batch(cmds...)
	}

	//nolint:gocritic
	switch msgType := msg.(type) {
	case tea.KeyMsg:
		switch msgType.String() {
		case "enter":
			m.selectedItem = m.list.SelectedItem()
		case "esc":
			m.textInput.Reset()
			m.list.ResetSelected()
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View is the Bubble Text text renderer.
func (m Model) View() string {
	// Set the paginator options with the current page and total pages from
	// the list component.
	po := Paginator{
		Position: m.list.Paginator.Page,
		Total:    m.list.Paginator.TotalPages,
		Height:   m.Height,
		Styles:   m.Paginator.Styles,
	}

	// Join the text input and list components vertically.
	left := lipgloss.JoinVertical(lipgloss.Top, m.textInput.View(), m.list.View())
	right := NewPaginator(po)

	// Join the text input and list components to the paginator.
	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

// Focused return the focus state of the model.
func (m Model) Focused() bool {
	return m.focus
}

// Focus sets the focus state of the model. When the model
// is in focus it can receive keyboard input.
func (m *Model) Focus() {
	m.focus = true
}

// Blur removes the focus state of the model. When the model
// is blurred it cannot receive keyboard input.
func (m *Model) Blur() {
	m.focus = false
}

// setModels sets the style of the components of the model.
func (m *Model) setModels() tea.Cmd {
	// Merge the default styles with any overrides.
	m.List.Styles = mergeListStyles(m.List.Styles)
	m.TextInput.Styles = mergeTextInputStyles(m.TextInput.Styles)
	m.Paginator.Styles = mergePaginatorStyles(m.Paginator.Styles)

	// Set the base state of the list and text input.
	m.setList()
	m.setTextInput()

	// Match the text input focus with the focus state of the model.
	switch {
	case m.focus && !m.textInput.Focused():
		return m.textInput.Focus()
	case !m.focus && m.textInput.Focused():
		m.textInput.Blur()
	}

	return nil
}
