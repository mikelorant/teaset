package filterlist

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// List is the List widget model.
type List struct {
	Width  int
	Height int
	Styles ListStyles
}

// ListStyles is the styling of the list widget.
type ListStyles struct {
	// Item style.
	Item lipgloss.Style

	// The item prompt is a left border style.
	ItemPrompt lipgloss.Border

	// The selected item.
	ItemSelected lipgloss.Style

	// The item prompt indicator character.
	ItemIndicator string

	// Style of empty list.
	NoItems lipgloss.Style
}

const (
	defaultItemIndicator = "❯"
)

// NewList creates a new list widget.
func NewList(lo List) list.Model {
	l := list.New(nil, NewDelegate(mergeListStyles(lo.Styles)), lo.Width, lo.Height)
	l.SetShowHelp(false)
	l.SetShowPagination(false)
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)
	l.SetShowFilter(false)

	return l
}

// NewDelegate creates a new list delegate. This is a modified default delegate with
// spacing and descriton disabled to make the output compact and single line.
func NewDelegate(styles ListStyles) list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	d.SetSpacing(0)
	d.ShowDescription = false
	d.Styles = NewDefaultItemStyles(styles)

	return d
}

// NewDefaultItemStyles sets the style of both the item and selected item.
func NewDefaultItemStyles(styles ListStyles) list.DefaultItemStyles {
	s := list.NewDefaultItemStyles()
	s.NormalTitle = styles.Item
	s.SelectedTitle = styles.ItemSelected

	return s
}

// SetItems set the items in the list.
func (m *Model) SetItems(is []list.Item) tea.Cmd {
	return m.list.SetItems(is)
}

// Selected item selects the current item.
//
//nolint:ireturn
func (m Model) SelectedItem() list.Item {
	return m.list.SelectedItem()
}

// Select moves the selected item to the index specified.
func (m *Model) Select(i int) {
	m.list.Select(i)
}

// setList sets the list dimension and styles.
func (m *Model) setList() {
	// Remove selected indicator (1) + selected padding (1) +
	// margin of paginator (1) and paginator (1) = 4
	m.list.SetWidth(m.Width - 4)
	// Text input uses the first line.
	m.list.SetHeight(m.Height - 1)
	m.list.SetDelegate(NewDelegate(m.List.Styles))
}

// ToItems casts the list of items so they ca be used with
// the list component.
func ToItems[T list.Item](v []T) []list.Item {
	items := make([]list.Item, len(v))
	for idx, i := range v {
		items[idx] = i
	}

	return items
}

// mergeListStyles merges the default styles with any existing
// defined styles.
func mergeListStyles(ls ListStyles) ListStyles {
	bs := lipgloss.Border{Left: ls.ItemIndicator}

	ls.Item = ls.Item.PaddingLeft(2)
	ls.ItemSelected = ls.ItemSelected.BorderStyle(bs).BorderLeft(true).PaddingLeft(1)

	return ls
}
