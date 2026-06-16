public class Solution {
    public bool IsValidSudoku(char[][] board) {
        for (int i = 0; i < board.Length; i++) {
            if (!IsRowValid(board[i]))
                return false;
        }

        for (int j = 0; j < board[0].Length; j++) {
            if (!IsColValid(board, j))
                return false;
        }

        return AreSquaresValid(board);
    }

    public bool IsRowValid(char[] row) {
        var set = new HashSet<char>();
        foreach (var c in row) {
            if (c == '.')
                continue;
            if (set.Contains(c))
                return false;
            set.Add(c);
        }
        return true;
    }

    public bool IsColValid(char[][] board, int colIdx) {
        char[] col = new char[board.Length];
        for (int i = 0; i < board.Length; i++) {
            col[i] = board[i][colIdx];
        }
        return IsRowValid(col);
    }

    public bool AreSquaresValid(char[][] board) {
        char[] square = new char[9];
        for (int i = 0; i < board.Length; i += 3) {
            for (int j = i; j < board[i].Length; j += 3) {
              square[0] = board[i][j];
              square[1] = board[i][j+1];
              square[2] = board[i][j+2];
              square[3] = board[i+1][j];
              square[4] = board[i+1][j+1];
              square[5] = board[i+1][j+2];
              square[6] = board[i+2][j];
              square[7] = board[i+2][j+1];
              square[8] = board[i+2][j+2];
              if(!IsRowValid(square)) return false;
            }
        }
        return true;
    }
}
