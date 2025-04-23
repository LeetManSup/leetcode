/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled
cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the
digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not
necessarily solvable.
Only the filled cells need to be validated according to the
mentioned rules.
*/

func isValidSudoku(board [][]byte) bool {
	validateNumPack := func(elements ...byte) bool {
		numSet := make(map[byte]bool)
		for _, el := range elements {
			if el >= byte('1') && el <= byte('9') {
				if numSet[el] {
					return false
				}
				numSet[el] = true
			}
		}
		return true
	}

	for col := range len(board) {
		column := []byte{}
		for row := range len(board) {
			if !validateNumPack(board[row]...) {
				return false
			}
			column = append(column, board[row][col])
		}
		if !validateNumPack(column...) {
			return false
		}
	}

	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board); j += 3 {
			subBoxNums := []byte{}
			subBoxNums = append(subBoxNums, board[i][j:j+3]...)
			subBoxNums = append(subBoxNums, board[i+1][j:j+3]...)
			subBoxNums = append(subBoxNums, board[i+2][j:j+3]...)
			if !validateNumPack(subBoxNums...) {
				return false
			}
		}
	}

	return true
}