package filterlist

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

// Text input is the text input widget model.
type TextInput struct {
	Width int

	// The leading character for the prompt text.
	PromptMark string

	// The text before any user input.
	PromptText string

	// Placeholder text that is cleared once there is some user input.
	Placeholder string

	// Character limit of the filter text input.
	CharLimit int
	Styles    TextInputStyles
}

type TextInputStyles struct {
	PromptMark  lipgloss.Style
	PromptText  lipgloss.Style
	Text        lipgloss.Style
	Placeholder lipgloss.Style
	Cursor      lipgloss.Style
}

const (
	defaultPromptMark = "?"
	defaultPromptText = "Filter:"
)

// NewTextInput creates a new text input widget.
func NewTextInput(tio TextInput) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = tio.Placeholder
	ti.CharLimit = tio.CharLimit
	ti.Prompt = textInputPrompt(tio)
	ti.Width = tio.Width - lipgloss.Width(ti.Prompt)

	return ti
}

// Filter returns the text entered into the text input.
func (m Model) Filter() string {
	return m.textInput.Value()
}

// setTextInput sets the default state of the text input.
func (m *Model) setTextInput() {
	m.textInput.Prompt = textInputPrompt(m.TextInput)
	// Text input width is calculated excluding the prompt.
	// Remove selected indicator (1) + selected padding (1) +
	// margin of paginator (1) and paginator (1) = 4
	m.textInput.Width = m.Width - lipgloss.Width(m.textInput.Prompt) - 4
	m.textInput.Placeholder = m.TextInput.Placeholder
	m.textInput.TextStyle = m.TextInput.Styles.Text
	m.textInput.Prompt = textInputPrompt(m.TextInput)
}

// textInputPrompt merges the prompt mark and prompt text together.
func textInputPrompt(tio TextInput) string {
	pm := tio.Styles.PromptMark.Render(tio.PromptMark)
	pt := tio.Styles.PromptText.Render(tio.PromptText)

	return lipgloss.JoinHorizontal(lipgloss.Top, pm, pt)
}

// mergeListStyles merges the default styles with any existing
// defined styles.
func mergeTextInputStyles(tis TextInputStyles) TextInputStyles {
	tis.PromptMark = tis.PromptMark.MarginRight(1)
	tis.PromptText = tis.PromptText.MarginRight(1)

	return tis
}
