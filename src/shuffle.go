package sudoku

import "math/rand"

func ShuffleArr[T any](arr [9]T) [9]T {
	result := arr
	sliceView := result[:]
	rand.Shuffle(len(sliceView), func(i, j int) {
		sliceView[i], sliceView[j] = sliceView[j], sliceView[i]
	})
	return result
}
