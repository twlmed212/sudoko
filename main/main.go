package main

import (
	"fmt"
	"os"
)

const N = 9

func printBoard(board [][]byte) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func isSafe(board [][]byte, row, col int, num byte) bool {
	// Check if the number is already present in the row or column
	for i := 0; i < N; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// Check if the number is already present in the 3x3 subgrid
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

func solveSudoku(board [][]byte) bool {
	var row, col int
	var foundEmpty bool

	// Find the first empty cell
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == '.' {
				row, col = i, j
				foundEmpty = true
				break
			}
		}
		if foundEmpty {
			break
		}
	}

	// If no empty cell is found, Sudoku is solved
	if !foundEmpty {
		return true
	}

	// Try placing numbers from '1' to '9'
	for num := byte('1'); num <= '9'; num++ {
		if isSafe(board, row, col, num) {
			// Place the number if it's safe
			board[row][col] = num

			// Recur to solve the rest of the Sudoku
			if solveSudoku(board) {
				return true
			}

			// If placing the number doesn't lead to a solution, backtrack
			board[row][col] = '.'
		}
	}

	// If no number can be placed, the puzzle is unsolvable
	return false
}

func isValidSudoku(board [][]byte) bool {
	// Check rows
	for i := 0; i < 9; i++ {
		seen := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if seen[board[i][j]] {
					return false
				}
				seen[board[i][j]] = true
			}
		}
	}

	// Check columns
	for j := 0; j < 9; j++ {
		seen := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' {
				if seen[board[i][j]] {
					return false
				}
				seen[board[i][j]] = true
			}
		}
	}

	// Check 3x3 subgrids
	for block := 0; block < 9; block++ {
		seen := make(map[byte]bool)
		startRow := (block / 3) * 3
		startCol := (block % 3) * 3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[startRow+i][startCol+j] != '.' {
					if seen[board[startRow+i][startCol+j]] {
						return false
					}
					seen[board[startRow+i][startCol+j]] = true
				}
			}
		}
	}

	return true
}

func main() {
	args := os.Args[1:]

	// Validate the number of arguments
	if len(args) != N {
		fmt.Println("Error")
		return
	}

	board := make([][]byte, N)
	for i := range board {
		// Validate the length of each row
		if len(args[i]) != N {
			fmt.Println("Error")
			return
		}

		// Validate characters in each row
		for _, char := range args[i] {
			if char != '.' && (char < '1' || char > '9') {
				fmt.Println("Error")
				return
			}
		}

		board[i] = []byte(args[i])
	}

	// Validate Sudoku boardfmt.Println("Input Sudoku:")
	// printBoard(board)
	if !isValidSudoku(board) {
		fmt.Println("Error")
		return
	}

	// fmt.Println("Input Sudoku:")
	// printBoard(board)

	// Solve the Sudoku puzzle
	if solveSudoku(board) {
		//	fmt.Println("Solved Sudoku:")
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}
