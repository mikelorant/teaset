package toggle

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	Inactive = "▢"
	Active   = "▣"
)

type Styles struct {
	HeadingActive   lipgloss.Style
	HeadingInactive lipgloss.Style
	Text            lipgloss.Style
	Active          lipgloss.Style
	Inactive        lipgloss.Style
}

func defaultStyles() Styles {
	var s Styles

	// Focused component heading.
	s.HeadingActive = lipgloss.NewStyle().Bold(true)

	// Blurred component heading.
	s.HeadingInactive = lipgloss.NewStyle()

	// Enabled toggle switch.
	s.Active = lipgloss.NewStyle().SetString(Active)

	// Disabled toggle switch.
	s.Inactive = lipgloss.NewStyle().SetString(Inactive)

	// Style of switch text.
	s.Text = lipgloss.NewStyle()

	return s
}
