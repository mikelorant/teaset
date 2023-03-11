package radio_test

import (
	"testing"

	"github.com/mikelorant/teaset/radio"
	"github.com/mikelorant/teaset/uitest"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		heading string
		values  []string
		loop    bool
		model   func(m radio.Model) radio.Model
	}

	type want struct {
		model func(m radio.Model)
	}

	tests := map[string]struct {
		name string
		args args
		want want
	}{
		"default": {
			args: args{
				heading: "test",
				values:  []string{"1", "2", "3", "4", "5"},
			},
			want: want{
				model: func(m radio.Model) {
					assert.Equal(t, 0, m.Index(), "Index")
					assert.Equal(t, "1", m.Value(), "Value")
				},
			},
		},
		"empt": {
			want: want{
				model: func(m radio.Model) {
					assert.Empty(t, m.Value())
				},
			},
		},
		"next": {
			args: args{
				heading: "test",
				values:  []string{"1", "2", "3", "4", "5"},
				model: func(m radio.Model) radio.Model {
					m.Next()
					m.Next()

					return m
				},
			},
			want: want{
				model: func(m radio.Model) {
					assert.Equal(t, 2, m.Index(), "Index")
					assert.Equal(t, "3", m.Value(), "Value")
				},
			},
		},
		"select": {
			args: args{
				heading: "test",
				values:  []string{"1", "2", "3", "4", "5"},
				model: func(m radio.Model) radio.Model {
					m.Select("3")

					return m
				},
			},
			want: want{
				model: func(m radio.Model) {
					assert.Equal(t, 2, m.Index(), "Index")
					assert.Equal(t, "3", m.Value(), "Value")
				},
			},
		},
		"previous": {
			args: args{
				heading: "test",
				values:  []string{"1", "2", "3", "4", "5"},
				model: func(m radio.Model) radio.Model {
					m.Select("5")
					m.Previous()

					return m
				},
			},
			want: want{
				model: func(m radio.Model) {
					assert.Equal(t, 3, m.Index(), "Index")
					assert.Equal(t, "4", m.Value(), "Value")
				},
			},
		},
		"forward_boundary": {
			args: args{
				heading: "test",
				values:  []string{"1", "2", "3", "4", "5"},
				model: func(m radio.Model) radio.Model {
					m.Select("5")
					m.Next()

					return m
				},
			},
			want: want{
				model: func(m radio.Model) {
					assert.Equal(t, 4, m.Index(), "Index")
					assert.Equal(t, "5", m.Value(), "Value")
				},
			},
		},
		"forward_boundary_loop": {
			args: args{
				heading: "test",
				values:  []string{"1", "2", "3", "4", "5"},
				loop:    true,
				model: func(m radio.Model) radio.Model {
					m.Select("5")
					m.Next()

					return m
				},
			},
			want: want{
				model: func(m radio.Model) {
					assert.Equal(t, 0, m.Index(), "Index")
					assert.Equal(t, "1", m.Value(), "Value")
				},
			},
		},
		"backward_boundary": {
			args: args{
				heading: "test",
				values:  []string{"1", "2", "3", "4", "5"},
				model: func(m radio.Model) radio.Model {
					m.Previous()

					return m
				},
			},
			want: want{
				model: func(m radio.Model) {
					assert.Equal(t, 0, m.Index(), "Index")
					assert.Equal(t, "1", m.Value(), "Value")
				},
			},
		},
		"backward_boundary_loop": {
			args: args{
				heading: "test",
				values:  []string{"1", "2", "3", "4", "5"},
				loop:    true,
				model: func(m radio.Model) radio.Model {
					m.Previous()

					return m
				},
			},
			want: want{
				model: func(m radio.Model) {
					assert.Equal(t, 4, m.Index(), "Index")
					assert.Equal(t, "5", m.Value(), "Value")
				},
			},
		},
		"down": {
			args: args{
				heading: "test",
				values:  []string{"1", "2", "3", "4", "5"},
				model: func(m radio.Model) radio.Model {
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})

					return m
				},
			},
			want: want{
				model: func(m radio.Model) {
					assert.Equal(t, 2, m.Index(), "Index")
					assert.Equal(t, "3", m.Value(), "Value")
				},
			},
		},
		"up": {
			args: args{
				heading: "test",
				values:  []string{"1", "2", "3", "4", "5"},
				model: func(m radio.Model) radio.Model {
					m.Focus()

					m.Select("5")

					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyUp})
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyUp})

					return m
				},
			},
			want: want{
				model: func(m radio.Model) {
					assert.Equal(t, 2, m.Index(), "Index")
					assert.Equal(t, "3", m.Value(), "Value")
				},
			},
		},
		"focus": {
			args: args{
				values: []string{"1", "2", "3", "4", "5"},
				model: func(m radio.Model) radio.Model {
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})

					return m
				},
			},
			want: want{
				model: func(m radio.Model) {
					assert.True(t, m.Focused())
					assert.Equal(t, 1, m.Index(), "Index")
					assert.Equal(t, "2", m.Value(), "Value")
				},
			},
		},
		"blur": {
			args: args{
				values: []string{"1", "2", "3", "4", "5"},
				model: func(m radio.Model) radio.Model {
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
					m.Blur()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})

					return m
				},
			},
			want: want{
				model: func(m radio.Model) {
					assert.False(t, m.Focused())
					assert.Equal(t, 1, m.Index(), "Index")
					assert.Equal(t, "2", m.Value(), "Value")
				},
			},
		},
	}

	for name, tt := range tests {
		name := name
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			m := radio.New()
			m.Heading = tt.args.heading
			m.Values = tt.args.values
			m.Loop = tt.args.loop

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
