'''
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
'''

class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        def validate(*args: str) -> bool:
            visited = set()
            for arg in args:
                if arg == ".":
                    continue
                if arg in visited:
                    return False
                visited.add(arg)
            return True

        n = len(board)

        for i in range(n):
            if not validate(*board[i]):
                return False

            column = [board[j][i] for j in range(n)]
            if not validate(*column):
                return False

        for i in range(0, n, 3):
            for j in range(0, n, 3):
                subbox = [
                    board[i + x][j + y]
                    for x in range(3)
                    for y in range(3)
                ]
                if not validate(*subbox):
                    return False

        return True