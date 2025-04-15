package sudoku

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Sudoku struct {
	Grid        *Grid
	Subsections Subsections
	Solution    Grid
}

type Candidates [9]uint8

type Cell struct {
	Value      uint8
	Candidates [9]uint8
}

type Row [9]Cell
type Column [9]Cell
type Box [9]Cell

type Grid [9][9]Cell

type Subsection [9]*Cell

type Subsections struct {
	Rows, Columns, Boxes [9]Subsection
}

func (g *Grid) String() string {
	var b strings.Builder
	for _, row := range g {
		b.WriteRune('|')
		for i := range row {
			if row[i].Value > 0 {
				yellowStr := color.YellowString("%d", row[i].Value)
				b.WriteString(fmt.Sprintf("%s|", yellowStr))
			} else {
				b.WriteString(fmt.Sprintf("%d|", row[i].Value))
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (s *Subsection) String() string {
	var b strings.Builder
	b.WriteString("[")
	for i, cell := range s {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(fmt.Sprintf("%d", cell.Value))
	}
	b.WriteString("]")
	return b.String()
}
