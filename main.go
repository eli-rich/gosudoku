package main

import (
	_ "embed"
	"fmt"
	"os"

	sudoku "github.com/eli-rich/gosudoku/src"
)

//go:embed puzzle.txt
var puzzle string

func main() {
	sPuzzle, err := sudoku.NewSudoku(puzzle)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	uniq := sPuzzle.Solve(true, false)
	fmt.Println(sPuzzle.Solution.String())
	if uniq {
		fmt.Println("This puzzle is unique!")
	} else {
		fmt.Println("This puzzle is not unique!")
	}
}
