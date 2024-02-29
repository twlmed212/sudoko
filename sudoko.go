package piscine

import "fmt"

const N = 9

func printGrid(grid [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
}

func isValid(grid [N][N]int, row, col, num int) bool {
	for x := 0; x < N; x++ {
		if grid[row][x] == num || grid[x][col] == num {
			return false
		}
	}
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

func solveSudoku(grid [N][N]int) bool {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if grid[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(grid, row, col, num) {
						grid[row][col] = num
						if solveSudoku(grid) {
							return true
						}
						grid[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}
