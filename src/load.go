package sudoku

import (
	"fmt"
	"strconv"
	"strings"
)

func GenerateGrid(puzzle string) (*Grid, error) {
	puzzle = strings.TrimSpace(puzzle)
	grid := &Grid{}
	row, col := 0, 0
	for _, char := range puzzle {
		if char == '\n' {
			row++
			col = 0
			continue
		}
		if col >= 9 {
			col = 0
			row++
		}
		if row >= 9 {
			return nil, fmt.Errorf("puzzle has too many cells")
		}
		value, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, fmt.Errorf("invalid char in puzzle: %c", char)
		}
		grid[row][col] = Cell{Value: uint8(value)}
		col++
	}
	return grid, nil
}

func (grid *Grid) GenerateSubsections() Subsections {
	subsections := Subsections{}
	// row pointers
	for r := range 9 {
		for c := range 9 {
			subsections.Rows[r][c] = &grid[r][c]
			subsections.Columns[c][r] = &grid[r][c]
			boxIdx := 3*(r/3) + (c / 3)
			boxPos := 3*(r%3) + (c % 3)
			subsections.Boxes[boxIdx][boxPos] = &grid[r][c]
		}
	}
	return subsections
}

func NewSudoku(puzzle string) (*Sudoku, error) {
	grid, err := GenerateGrid(puzzle)
	if err != nil {
		return nil, err
	}
	return &Sudoku{
		Grid:        grid,
		Subsections: grid.GenerateSubsections(),
	}, nil
}
