package sudoku

func (s *Sudoku) getCandidates(row, column int, generating bool) Candidates {
	valid := [9]bool{true, true, true, true, true, true, true, true, true}
	boxIdx := 3*(row/3) + (column / 3)

	candidates := Candidates{}

	for i := range 9 {
		if cellValue := s.Subsections.Rows[row][i].Value; cellValue > 0 {
			valid[cellValue-1] = false
		}
		if cellValue := s.Subsections.Columns[column][i].Value; cellValue > 0 {
			valid[cellValue-1] = false
		}
		if cellValue := s.Subsections.Boxes[boxIdx][i].Value; cellValue > 0 {
			valid[cellValue-1] = false
		}
	}
	for i := range 9 {
		if valid[i] {
			candidates[i] = uint8(i + 1)
		}
	}
	if generating && len(candidates) > 0 {
		candidates = ShuffleArr(candidates)
	}
	return candidates
}

func (s *Sudoku) checkCandidate(row, col int, candidate uint8) bool {
	candidates := s.getCandidates(row, col, false)
	for _, c := range candidates {
		if c == candidate {
			return true
		}
	}
	return false
}

func (s *Sudoku) GenerateCandidates(generating bool) {
	for i := range 9 {
		for j := range 9 {
			if s.Grid[i][j].Value == 0 {
				s.Grid[i][j].Candidates = s.getCandidates(i, j, generating)
			}
		}
	}
}
