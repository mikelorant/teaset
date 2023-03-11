package toggle

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Model is the Bubble Tea model for the Toggle widget.
type Model struct {
	// Heading is printed on top of the Toggle switch.
	Heading string

	// State is the state of the Toggle switch.
	State bool

	// Styles all the widget components.
	Styles Styles

	// Text is the label for the toggle.
	Text string

	// Focus is the state that determines whether
	// keyboard inputs should be accepted.
	focus bool
}

const defaultText = "Enable"

// Instaniate a new Bubble Tea model.
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
			case " ":
				m.Toggle()
			}
		}
	}

	return m, nil
}

// View is the Bubble Text text renderer.
func (m Model) View() string {
	return m.renderToggle()
}

// Toggle switches the state.
func (m *Model) Toggle() {
	m.State = !m.State
}

// renderToggle rends the toggle widget.
func (m Model) renderToggle() string {
	var arr []string

	// If empty heading, do not render it.
	if m.Heading != "" {
		heading := m.Styles.HeadingInactive.Render(m.Heading)
		// Use a highlighted header if component has focus.
		if m.focus {
			heading = m.Styles.HeadingActive.Render(m.Heading)
		}

		arr = append(arr, heading)
	}

	state := m.Styles.Inactive.String()
	if m.State {
		state = m.Styles.Active.String()
	}

	text := defaultText
	if m.Text != "" {
		text = m.Text
	}

	arr = append(arr, fmt.Sprintf("%v %v", state, m.Styles.Text.Render(text)))

	return strings.Join(arr, "\n")
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
