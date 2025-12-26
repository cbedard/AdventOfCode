package solutions2025

import (
	"fmt"
	"strconv"
	"strings"
)

func Day6Part1(lines []string) {
	total := 0
	numberGrid := [][]int{}

	for i := 0; i < len(lines)-1; i++ {
		numberGrid = append(numberGrid, parseInts(lines[i]))
	}

	operators := strings.Fields(lines[len(lines)-1])
	for i := range operators {
		var curr int
		if operators[i] == "*" {
			curr = 1
			for j := range numberGrid {
				curr *= numberGrid[j][i]
			}
		} else {
			curr = 0
			for j := range numberGrid {
				curr += numberGrid[j][i]
			}
		}
		total += curr
	}

	fmt.Println(total)
}

func Day6Part2(lines []string) {
	total := 0

	operators := strings.Fields(lines[len(lines)-1])
	low, high := 0, 1

	for i := range operators {
		var curr int
		for high < len(lines[0]) && !isColumnAllSpaces(lines, high) {
			high++
		}

		currNums := []int{}
		for low < high {
			currNums = append(currNums, parseColumnInt(lines[:len(lines)-1], low))
			low++
		}

		if operators[i] == "*" {
			curr = 1
			for j := range currNums {
				curr *= currNums[j]
			}
		} else {
			curr = 0
			for j := range currNums {
				curr += currNums[j]
			}
		}

		total += curr
		low = high + 1
		high = low + 1
	}

	fmt.Println(total)
}

func parseColumnInt(grid []string, col int) int {
	str := ""
	for i := range len(grid) {
		if grid[i][col] != ' ' {
			str += string(grid[i][col])
		}
	}

	num, _ := strconv.Atoi(str)
	return num
}

func isColumnAllSpaces(grid []string, col int) bool {
	for i := range len(grid) {
		if grid[i][col] != ' ' {
			return false
		}
	}
	return true
}

func parseInts(s string) []int {
	parts := strings.Fields(s)
	nums := make([]int, len(parts))
	for i, p := range parts {
		nums[i], _ = strconv.Atoi(p)
	}
	return nums
}
