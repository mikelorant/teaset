package filterlist_test

import (
	"testing"

	"github.com/mikelorant/teaset/filterlist"
	"github.com/mikelorant/teaset/uitest"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/assert"
)

type MockItem struct {
	title string
}

func (i MockItem) Title() string {
	return i.title
}

func (i MockItem) Description() string {
	return "Test"
}

func (i MockItem) FilterValue() string {
	return "Test"
}

func TestModel(t *testing.T) {
	t.Parallel()

	type args struct {
		model func(filterlist.Model) filterlist.Model
	}

	type want struct {
		model func(filterlist.Model)
	}

	tests := map[string]struct {
		args args
		want want
	}{
		"default": {},
		"one": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := []MockItem{
						{title: "item"},
					}
					m.SetItems(filterlist.ToItems(items))

					return m
				},
			},
		},
		"none": {},
		"multiple_pages": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))

					return m
				},
			},
		},
		"height": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.Height = 10
					m.SetItems(filterlist.ToItems(items))

					return m
				},
			},
		},
		"width": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.Width = 40
					m.SetItems(filterlist.ToItems(items))

					return m
				},
			},
		},
		"overflow": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.Height = 3
					m.SetItems(filterlist.ToItems(items))

					return m
				},
			},
		},
		"prompt_text": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.TextInput.PromptText = "test"

					return m
				},
			},
		},
		"focus": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()

					return m
				},
			},
			want: want{
				model: func(m filterlist.Model) {
					assert.True(t, m.Focused())
				},
			},
		},
		"blur": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					m.Focus()
					m, _ = m.Update(nil)
					m.Blur()

					return m
				},
			},
			want: want{
				model: func(m filterlist.Model) {
					assert.False(t, m.Focused())
				},
			},
		},
		"down": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})

					return m
				},
			},
		},
		"up": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyUp})

					return m
				},
			},
		},
		"pagedown": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgDown})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgDown})

					return m
				},
			},
		},
		"pagedown_lastpage": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgDown})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgDown})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgDown})

					return m
				},
			},
		},
		"pageup": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgDown})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgDown})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgUp})

					return m
				},
			},
		},
		"pageup_firstpage": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgUp})

					return m
				},
			},
		},
		"enter": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})

					return m
				},
			},
			want: want{
				model: func(m filterlist.Model) {
					title := m.SelectedItem().(MockItem).title
					assert.Equal(t, "item 2345", title)
				},
			},
		},
		"type": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Focus()
					m = sendString(m, "test")

					return m
				},
			},
		},
		"escape": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Focus()
					m = sendString(m, "test")
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyEsc})

					return m
				},
			},
		},
		"set_items": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Focus()

					m.SetItems(filterlist.ToItems([]MockItem{
						{title: "item 4321"},
					}))

					return m
				},
			},
		},
		"filter": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Focus()
					m = sendString(m, "test")

					return m
				},
			},
			want: want{
				model: func(m filterlist.Model) {
					assert.Equal(t, "test", m.Filter())
				},
			},
		},
		"select": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.Select(1)

					return m
				},
			},
			want: want{
				model: func(m filterlist.Model) {
					assert.Equal(t, "item 2345", m.SelectedItem().(MockItem).Title())
				},
			},
		},
		"override": {
			args: args{
				model: func(m filterlist.Model) filterlist.Model {
					items := testItems()
					m.SetItems(filterlist.ToItems(items))
					m.TextInput.Styles.PromptMark = lipgloss.NewStyle().MarginLeft(4)
					m.List.Styles.Item = lipgloss.NewStyle().MarginLeft(4)
					m.List.Styles.ItemSelected = lipgloss.NewStyle().MarginLeft(4)

					return m
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		name := name

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			t.Run(name, func(t *testing.T) {
				m := filterlist.New()

				if tt.args.model != nil {
					m = tt.args.model(m)
				}

				m, _ = m.Update(nil)

				if tt.want.model != nil {
					tt.want.model(m)
				}

				v := uitest.StripString(m.View())
				autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(name))
			})
		})
	}
}

func testItems() []MockItem {
	return []MockItem{
		{title: "item 1234"},
		{title: "item 2345"},
		{title: "item 3456"},
		{title: "item 4567"},
		{title: "item 5678"},
		{title: "item 6789"},
		{title: "item 7890"},
		{title: "item 8901"},
		{title: "item 9012"},
	}
}

//nolint:ireturn
func sendString(m filterlist.Model, str string) filterlist.Model {
	for _, r := range str {
		msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}}
		m, _ = m.Update(msg)
	}

	return m
}
