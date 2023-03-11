package radio

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Model is the Bubble Tea model for the Radio widget.
type Model struct {
	// Heading is printed on top of the Radio selector.
	Heading string

	// Values are all the possible values selectable
	// with the Radio.
	Values []string

	// Loop determines whether the boundary loops
	// back to the first or last value.
	Loop bool

	// Styles all the widget components.
	Styles Styles

	// Index is the index of all possible values.
	index int

	// Focus is the state that determines whether
	// keyboard inputs should be accepted.
	focus bool
}

// New creates a new model with default styles.
func New() Model {
	return Model{
		Styles: defaultStyles(),
	}
}

// Update is the Bubble Tea update loop.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if m.focus {
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "down":
				m.Next()
			case "up":
				m.Previous()
			}
		}
	}

	return m, nil
}

// View is the Bubble Text text renderer.
func (m Model) View() string {
	return m.renderRadio()
}

// Value is the value currently selected.
func (m Model) Value() string {
	if len(m.Values) == 0 {
		return ""
	}

	return m.Values[m.index]
}

// Index is the currently selected index.
func (m Model) Index() int {
	return m.index
}

// renderRadio rends the radio widget.
func (m Model) renderRadio() string {
	var arr []string

	// If the heading is not empty we add it.
	if m.Heading != "" {
		heading := m.Styles.HeadingInactive.Render(m.Heading)
		if m.focus {
			heading = m.Styles.HeadingActive.Render(m.Heading)
		}

		arr = append(arr, heading)
	}

	for idx, val := range m.Values {
		state := m.Styles.Inactive.String()

		// If the index matches the current value
		// mark the value as active and render with
		// the active element.
		if idx == m.index {
			state = m.Styles.Active.String()
		}

		arr = append(arr, fmt.Sprintf("%v %v", state, m.Styles.Values.Render(val)))
	}

	return strings.Join(arr, "\n")
}

// Next moves the widget to the next value.
func (m *Model) Next() {
	switch {
	// If loop is enabled and the we are at the
	// end of the list, we reset to the first item.
	case m.index == (len(m.Values)-1) && m.Loop:
		m.index = 0

		return

	// If we are at the end of the index, do nothing.
	case m.index == len(m.Values)-1:
		return
	}

	// Increment the index.
	m.index++
}

// Previous moves the widget to the previous value.
func (m *Model) Previous() {
	switch {
	// If loop is enabled and we are at the
	// start of the list, we jump to the list item.
	case m.index == 0 && m.Loop:
		m.index = len(m.Values) - 1

		return

	// If we are at the beginning of the index, do nothing.
	case m.index == 0:
		return
	}

	// Reduce the index.
	m.index--
}

// Select sets the index to the heading provided.
func (m *Model) Select(val string) {
	for idx, v := range m.Values {
		if v == val {
			m.index = idx

			break
		}
	}
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
