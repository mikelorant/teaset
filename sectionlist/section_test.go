package sectionlist_test

import (
	"testing"

	"github.com/mikelorant/teaset/sectionlist"
	"github.com/mikelorant/teaset/uitest"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		sections sectionlist.Sections
		model    func(sectionlist.Model) sectionlist.Model
	}

	type want struct {
		model func(sectionlist.Model)
	}

	tests := map[string]struct {
		name string
		args args
		want want
	}{
		"default": {},
		"normal": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test", Values: []string{"1", "2", "3", "4", "5"}},
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test", m.Section())
					assert.Equal(t, "1", m.Value())
				},
			},
		},
		"section_first_value_first_next": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test", Values: []string{"1", "2", "3", "4", "5"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.Next()

					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test", m.Section())
					assert.Equal(t, "2", m.Value())
				},
			},
		},
		"section_first_value_last_next": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test", Values: []string{"1", "2", "3", "4", "5"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.ValueIndex(4)
					m.Next()

					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test", m.Section())
					assert.Equal(t, "5", m.Value())
				},
			},
		},
		"section_first_next_section": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test1", Values: []string{"1", "2", "3", "4", "5"}},
					{Name: "test2", Values: []string{"one", "two", "three", "four", "five"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.ValueIndex(4)
					m.Next()

					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test2", m.Section())
					assert.Equal(t, "one", m.Value())
				},
			},
		},
		"section_last_value_last_next": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test1", Values: []string{"1", "2", "3", "4", "5"}},
					{Name: "test2", Values: []string{"one", "two", "three", "four", "five"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.SectionIndex(1)
					m.ValueIndex(4)
					m.Next()

					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test2", m.Section())
					assert.Equal(t, "five", m.Value())
				},
			},
		},
		"section_first_value_second_previous": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test", Values: []string{"1", "2", "3", "4", "5"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.ValueIndex(1)
					m.Previous()

					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test", m.Section())
					assert.Equal(t, "1", m.Value())
				},
			},
		},
		"section_second_value_first_previous": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test1", Values: []string{"1", "2", "3", "4", "5"}},
					{Name: "test2", Values: []string{"one", "two", "three", "four", "five"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.SectionIndex(1)
					m.Previous()

					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test1", m.Section())
					assert.Equal(t, "5", m.Value())
				},
			},
		},
		"section_first_value_first_previous": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test1", Values: []string{"1", "2", "3", "4", "5"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.Previous()

					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test1", m.Section())
					assert.Equal(t, "1", m.Value())
				},
			},
		},
		"section_no_value": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test", Values: []string{}},
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test", m.Section())
					assert.Equal(t, "", m.Value())
				},
			},
		},
		"sections_with_one_section_no_value": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test1", Values: []string{"1", "2", "3", "4", "5"}},
					{Name: "test2", Values: []string{}},
					{Name: "test3", Values: []string{"one", "two", "three", "four", "five"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.SectionIndex(1)

					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test2", m.Section())
					assert.Equal(t, "", m.Value())
				},
			},
		},
		"set_value_invalid": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test", Values: []string{"1", "2", "3", "4", "5"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.ValueIndex(5)

					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test", m.Section())
					assert.Equal(t, "1", m.Value())
				},
			},
		},
		"set_section_invalid": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test", Values: []string{"1", "2", "3", "4", "5"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.SectionIndex(1)

					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.Equal(t, "test", m.Section())
					assert.Equal(t, "1", m.Value())
				},
			},
		},
		"down": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test", Values: []string{"1", "2", "3", "4", "5"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.True(t, m.Focused())
					assert.Equal(t, "test", m.Section())
					assert.Equal(t, "3", m.Value())
				},
			},
		},
		"up": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test", Values: []string{"1", "2", "3", "4", "5"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.Focus()
					m.ValueIndex(4)
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyUp})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyUp})
					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.True(t, m.Focused())
					assert.Equal(t, "test", m.Section())
					assert.Equal(t, "3", m.Value())
				},
			},
		},
		"focus": {
			args: args{
				sections: sectionlist.Sections{
					{Name: "test", Values: []string{"1", "2", "3", "4", "5"}},
				},
				model: func(m sectionlist.Model) sectionlist.Model {
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
					m.Blur()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyUp})
					return m
				},
			},
			want: want{
				model: func(m sectionlist.Model) {
					assert.False(t, m.Focused())
					assert.Equal(t, "test", m.Section())
					assert.Equal(t, "2", m.Value())
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		name := name

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			m := sectionlist.New()
			m.Sections = tt.args.sections

			if tt.args.model != nil {
				m = tt.args.model(m)
			}

			if tt.want.model != nil {
				tt.want.model(m)
			}

			v := uitest.StripString(m.View())
			autogold.ExpectFile(t, autogold.Raw(v), autogold.Name(name))
		})
	}
}
