package sectionlist

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Sections []Section

type Section struct {
	Name   string
	Values []string
}

// Model is the Bubble Tea model for the Section widget.
type Model struct {
	Header   string
	Sections Sections
	Styles   Styles

	focus        bool
	sectionIndex int
	valueIndex   int
}

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
	return m.renderSectionList()
}

// Section returns the value for the current section.
func (m Model) Section() string {
	return m.Sections[m.sectionIndex].Name
}

// Value returns the value for the currently selected item.
func (m Model) Value() string {
	if len(m.Sections[m.sectionIndex].Values) == 0 {
		return ""
	}

	return m.Sections[m.sectionIndex].Values[m.valueIndex]
}

// SectionIndex sets the index of the section.
func (m *Model) SectionIndex(n int) {
	if n >= len(m.Sections) {
		return
	}

	m.sectionIndex = n
}

// ValueIndex sets the index of the current section value.
func (m *Model) ValueIndex(n int) {
	if n >= len(m.Sections[m.sectionIndex].Values) {
		return
	}

	m.valueIndex = n
}

// renderSectionList renders all the sections with their values.
func (m Model) renderSectionList() string {
	var arr []string

	for _, section := range m.Sections {
		arr = append(arr, m.renderSection(section))
		if len(section.Values) != 0 {
			arr = append(arr, m.renderValues(section))
		}
		arr = append(arr, "")
	}

	return strings.TrimSpace(strings.Join(arr, "\n"))
}

// renderSection renders the section headings. If the section
// is selected a prompt is added.
func (m Model) renderSection(section Section) string {
	prompt := m.Styles.SectionEmpty.String()

	if m.isSection(section) {
		prompt = m.Styles.SectionSelected.String()
	}

	return fmt.Sprintf("%v %v", prompt, string(section.Name))
}

// renderValues renders the values. If the section is selected
// a prompt is added to the selected value. Values above the
// selected value have a joiner added.
func (m Model) renderValues(section Section) string {
	var arr []string

	// Skip if not the current selected section
	if !m.isSection(section) {
		for _, v := range section.Values {
			prompt := m.Styles.ValueEmpty.String()
			line := fmt.Sprintf("%v%v%v", strings.Repeat(" ", 2), prompt, v)
			arr = append(arr, line)
		}

		return strings.Join(arr, "\n")
	}

	var prompt string

	for idx, v := range section.Values {
		switch {
		// Selected value.
		case idx == m.valueIndex:
			prompt = m.Styles.ValueSelected.String()
		// Values before selected.
		case idx < m.valueIndex:
			prompt = m.Styles.ValueJoiner.String()
		// Values after selected.
		default:
			prompt = m.Styles.ValueEmpty.String()
		}

		line := fmt.Sprintf("%v%v%v", strings.Repeat(" ", 2), prompt, v)
		arr = append(arr, line)
	}

	return strings.Join(arr, "\n")
}

// isSection determines whether the section is selected.
func (m Model) isSection(section Section) bool {
	for idx, s := range m.Sections {
		if s.Name == section.Name && m.sectionIndex == idx {
			return true
		}
	}

	return false
}

func (m *Model) Next() {
	// Length of sections
	lastSectionIndex := len(m.Sections) - 1

	// Length of section values
	lastValueIndex := len(m.Sections[m.sectionIndex].Values) - 1

	switch {
	// Index at end of value but not on the last section
	case !(m.sectionIndex == lastSectionIndex) && m.valueIndex == lastValueIndex:
		m.sectionIndex++
		m.valueIndex = 0
	// Index at last value and last section
	case m.sectionIndex == lastSectionIndex && m.valueIndex == lastValueIndex:
	default:
		m.valueIndex++
	}
}

func (m *Model) Previous() {
	// Length of section values
	lastValueIndex := len(m.Sections[m.sectionIndex].Values) - 1

	switch {
	// Index not at first section and at first value
	case m.sectionIndex != 0 && m.valueIndex == 0:
		m.sectionIndex--
		m.valueIndex = lastValueIndex
	// Index at first section and first value
	case m.sectionIndex == 0 && m.valueIndex == 0:
	default:
		m.valueIndex--
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
