package day3

import "unicode"

func isSymbol(c byte) bool {
	return !unicode.IsDigit(rune(c)) && c != '.'
}

func isCloseToSymbol(grid []string, lineIdx int, digitIdx int) bool {
	result := false
	notUpperLine := lineIdx != 0
	notBottomLine := lineIdx != len(grid)-1
	notRightBorder := digitIdx != len(grid[lineIdx])-1
	notLeftBorder := digitIdx != 0

	if notUpperLine {
		result = result || isSymbol(grid[lineIdx-1][digitIdx])
		if notLeftBorder {
			result = result || isSymbol(grid[lineIdx-1][digitIdx-1])
		}
		if notRightBorder {
			result = result || isSymbol(grid[lineIdx-1][digitIdx+1])
		}
	}
	if notBottomLine {
		result = result || isSymbol(grid[lineIdx+1][digitIdx])
		if notLeftBorder {
			result = result || isSymbol(grid[lineIdx+1][digitIdx-1])
		}
		if notRightBorder {
			result = result || isSymbol(grid[lineIdx+1][digitIdx+1])
		}
	}
	if notRightBorder {
		result = result || isSymbol(grid[lineIdx][digitIdx+1])
	}
	if notLeftBorder {
		result = result || isSymbol(grid[lineIdx][digitIdx-1])
	}

	return result
}
