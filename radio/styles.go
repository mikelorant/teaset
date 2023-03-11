package radio

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	Inactive = "○"
	Active   = "●"
)

type Styles struct {
	HeadingActive   lipgloss.Style
	HeadingInactive lipgloss.Style
	Values          lipgloss.Style
	Active          lipgloss.Style
	Inactive        lipgloss.Style
}

func defaultStyles() Styles {
	var s Styles

	s.HeadingActive = lipgloss.NewStyle().Bold(true)
	s.HeadingInactive = lipgloss.NewStyle()
	s.Active = lipgloss.NewStyle().SetString(Active)
	s.Inactive = lipgloss.NewStyle().SetString(Inactive)
	s.Values = lipgloss.NewStyle()

	return s
}
