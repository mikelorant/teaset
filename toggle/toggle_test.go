package toggle_test

import (
	"testing"

	"github.com/mikelorant/teaset/toggle"
	"github.com/mikelorant/teaset/uitest"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/autogold/v2"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		heading string
		text    string
		state   bool
		model   func(m toggle.Model) toggle.Model
	}

	type want struct {
		model func(m toggle.Model)
	}

	tests := map[string]struct {
		name string
		args args
		want want
	}{
		"default": {},
		"heading": {
			args: args{
				heading: "heading",
			},
		},
		"text": {
			args: args{
				text: "text",
			},
		},
		"toggle": {
			args: args{
				model: func(m toggle.Model) toggle.Model {
					m.Focus()
					m.Toggle()

					return m
				},
			},
			want: want{
				model: func(m toggle.Model) {
					assert.True(t, m.Focused())
					assert.True(t, m.State)
				},
			},
		},
		"space": {
			args: args{
				model: func(m toggle.Model) toggle.Model {
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeySpace})

					return m
				},
			},
			want: want{
				model: func(m toggle.Model) {
					assert.True(t, m.State)
				},
			},
		},
		"space_blur": {
			args: args{
				model: func(m toggle.Model) toggle.Model {
					m.Focus()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeySpace})
					m.Blur()
					m, _ = m.Update(tea.KeyMsg{Type: tea.KeySpace})

					return m
				},
			},
			want: want{
				model: func(m toggle.Model) {
					assert.True(t, m.State)
				},
			},
		},
		"heading_inactive": {
			args: args{
				heading: "heading",
				model: func(m toggle.Model) toggle.Model {
					m, _ = m.Update(nil)

					return m
				},
			},
		},
		"heading_active": {
			args: args{
				heading: "heading",
				model: func(m toggle.Model) toggle.Model {
					m.Focus()
					m, _ = m.Update(nil)

					return m
				},
			},
			want: want{
				model: func(m toggle.Model) {
					assert.True(t, m.Focused())
				},
			},
		},
	}

	for name, tt := range tests {
		name := name
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			m := toggle.New()
			m.Heading = tt.args.heading
			m.Text = tt.args.text
			m.State = tt.args.state

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
