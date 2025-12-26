package solutions2025

import (
	"fmt"
)

var dirs = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func Day4Part1(lines []string) {
	grid := stringToByteArr(lines)
	total := 0

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '@' {
				neighborCount := neighbourCount(grid, i, j)

				if neighborCount < 4 {
					total++
				}
			}
		}
	}

	fmt.Println(total)
}

func Day4Part2(lines []string) {
	grid := stringToByteArr(lines)
	total := 0
	moveMade := true

	for moveMade {
		moveMade = false

		for i := range grid {
			for j := range grid[0] {
				if grid[i][j] == '@' {
					neighborCount := neighbourCount(grid, i, j)

					if neighborCount < 4 {
						total++
						moveMade = true
						grid[i][j] = '.'
					}
				}
			}
		}
	}

	fmt.Println(total)
}

func stringToByteArr(lines []string) [][]byte {
	grid := make([][]byte, len(lines))
	for i := range grid {
		grid[i] = []byte(lines[i])
	}
	return grid
}

func inBounds(n, m, i, j int) bool {
	if i >= 0 && i < n && j >= 0 && j < m {
		return true
	}
	return false
}

func neighbourCount(grid [][]byte, i, j int) int {
	neighbours := 0
	//iterate all 8 dirs
	for _, dir := range dirs {
		ni, nj := i+dir[0], j+dir[1]

		//in bounds
		if inBounds(len(grid), len(grid[0]), ni, nj) && grid[ni][nj] == '@' {
			neighbours++
		}
	}
	return neighbours
}
