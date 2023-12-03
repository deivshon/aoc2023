package day3

import "main/src/pkg/utils"

func getEndDigitIndex(grid []string, lineIdx int, middleDigitIdx int) int {
	i := middleDigitIdx
	for i < len(grid[lineIdx]) && utils.IsDigit(grid[lineIdx][i]) {
		i++
	}
	i--

	return i
}

func getStartDigitIndex(grid []string, lineIdx int, middleDigitIdx int) int {
	i := middleDigitIdx
	for i >= 0 && utils.IsDigit(grid[lineIdx][i]) {
		i--
	}
	i++

	return i
}

func getStartEndDigitIdx(grid []string, lineIdx int, middleDigitIdx int) (int, int) {
	return getStartDigitIndex(grid, lineIdx, middleDigitIdx), getEndDigitIndex(grid, lineIdx, middleDigitIdx)
}
