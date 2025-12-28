package solutions2025

import (
	"fmt"
	"strconv"
	"strings"
)

type Shape struct {
	Pattern   [3]string
	CellCount int
}

type Puzzle struct {
	Width, Height int
	Counts        []int
	TotalCells    int
}

func Day12Part1(lines []string) {
	_, puzzles, impossible := parse12(lines)
	fmt.Println(len(puzzles) - impossible)
}

func parse12(lines []string) (shapes []Shape, puzzles []Puzzle, impossible int) {
	i := 0

	// Parse 6 shapes (0-5), each is "N:" followed by 3 lines
	for range 6 {
		for lines[i] == "" || strings.HasSuffix(lines[i], ":") {
			i++
		}
		pattern := [3]string{lines[i], lines[i+1], lines[i+2]}
		cellCount := strings.Count(pattern[0]+pattern[1]+pattern[2], "#")
		shapes = append(shapes, Shape{pattern, cellCount})
		i += 3
	}

	// Parse puzzle lines: "WxH: c0 c1 c2 c3 c4 c5"
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		parts := strings.Split(lines[i], ": ")
		dims := strings.Split(parts[0], "x")
		w, _ := strconv.Atoi(dims[0])
		h, _ := strconv.Atoi(dims[1])

		counts := make([]int, 6)
		totalCells := 0
		for j, s := range strings.Fields(parts[1]) {
			counts[j], _ = strconv.Atoi(s)
			totalCells += counts[j] * shapes[j].CellCount
		}

		// impossible if more cells than grid area
		if totalCells > w*h {
			impossible++
		}
		puzzles = append(puzzles, Puzzle{w, h, counts, totalCells})
	}

	return
}
