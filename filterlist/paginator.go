package filterlist

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Paginator is the Paginator widget model.
type Paginator struct {
	Position int
	Total    int
	Height   int
	Styles   PaginatorStyles
}

type PaginatorStyles struct {
	// Paginator boundary style.
	Boundary lipgloss.Style

	// Indicator for a page.
	DotEmpty lipgloss.Style

	// Indicator for the current page.
	DotFilled lipgloss.Style
}

const (
	PaginatorDotEmpty  = "○"
	PaginatorDotFilled = "●"
)

// NewPaginator creates a new paginator list.
func NewPaginator(po Paginator) string {
	pos := po.Position
	total := po.Total
	height := po.Height

	if total == 0 {
		return ""
	}

	ps := make([]string, total)
	for p := range ps {
		ps[p] = po.Styles.DotEmpty.String()
	}

	ps[pos] = po.Styles.DotFilled.String()

	// If there are more pages than the height, we truncate the list.
	// This is not the best outcome but prevents overflow and rendering issues.
	if total >= po.Height {
		ps = ps[:height]
	}

	str := strings.Join(ps, "\n")

	return po.Styles.Boundary.Render(str)
}

// mergePaginatorStyles merges the default styles with any existing
// defined styles.
func mergePaginatorStyles(ps PaginatorStyles) PaginatorStyles {
	ps.Boundary = ps.Boundary.MarginLeft(1)
	ps.DotEmpty = ps.DotEmpty.SetString(PaginatorDotEmpty)
	ps.DotFilled = ps.DotFilled.SetString(PaginatorDotFilled)

	return ps
}
