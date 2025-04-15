package sudoku

func (s *Sudoku) Solve(verify, generating bool) bool {
	// collect empty cells
	var emptyCells [][2]int
	for r := range 9 {
		for c := range 9 {
			if s.Grid[r][c].Value == 0 {
				emptyCells = append(emptyCells, [2]int{r, c})
			}
		}
	}
	// generate initial candidates
	for _, pos := range emptyCells {
		r, c := pos[0], pos[1]
		candidates := s.getCandidates(r, c, generating)

		for i := range s.Grid[r][c].Candidates {
			s.Grid[r][c].Candidates[i] = 0 // reset
		}
		copy(s.Grid[r][c].Candidates[:], candidates[:])
	}

	idx := 0
	solutions := 0
	candidateIdx := make([]int, len(emptyCells))

	// backtracking loop
	for idx >= 0 && idx < len(emptyCells) {
		if verify && idx < 0 {
			return solutions == 1
		}
		r, c := emptyCells[idx][0], emptyCells[idx][1]
		cell := &s.Grid[r][c]

		found := false
		for candidateIdx[idx] < 9 && !found {
			candidate := cell.Candidates[candidateIdx[idx]]
			candidateIdx[idx]++
			if candidate == 0 {
				continue
			}
			if s.checkCandidate(r, c, candidate) {
				cell.Value = candidate
				found = true
			}
		}
		if !found {
			cell.Value = 0
			candidateIdx[idx] = 0
			idx--
			continue
		}
		idx++
		if verify && idx == len(emptyCells) {
			s.Solution = *s.Grid
			solutions++
			if solutions > 1 {
				return false
			}
			idx--
			r, c = emptyCells[idx][0], emptyCells[idx][1]
			s.Grid[r][c].Value = 0
		}
	}
	return solutions == 1 || !verify
}
