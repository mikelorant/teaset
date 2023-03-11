package sectionlist

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	SectionEmpty    = " "
	SectionSelected = "❯"

	ValueEmpty    = "  "
	ValueJoiner   = "│ "
	ValueSelected = "└▸"
)

type Styles struct {
	HeadingActive   lipgloss.Style
	HeadingInactive lipgloss.Style
	SectionEmpty    lipgloss.Style
	SectionSelected lipgloss.Style
	ValueEmpty      lipgloss.Style
	ValueSelected   lipgloss.Style
	ValueJoiner     lipgloss.Style
}

func defaultStyles() Styles {
	var s Styles

	s.HeadingActive = lipgloss.NewStyle().Bold(true)

	s.HeadingInactive = lipgloss.NewStyle()

	s.SectionEmpty = lipgloss.NewStyle().SetString(SectionEmpty)

	s.SectionSelected = lipgloss.NewStyle().SetString(SectionSelected)

	s.ValueEmpty = lipgloss.NewStyle().SetString(ValueEmpty)

	s.ValueJoiner = lipgloss.NewStyle().SetString(ValueJoiner)

	s.ValueSelected = lipgloss.NewStyle().SetString(ValueSelected)

	return s
}
