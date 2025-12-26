package solutions2025

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day5Part1(lines []string) {
	total := 0
	intervals, targets := parseInput(lines)

	for _, target := range targets {
		ok := slices.ContainsFunc(intervals, func(interval []int) bool {
			return target >= interval[0] && target <= interval[1]
		})

		if ok {
			total++
		}
	}

	fmt.Println(total)
}

func Day5Part2(lines []string) {
	total := 0
	intervals, _ := parseInput(lines)

	// sort by lower bound, ascending
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	compressed := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		// if our curr interval low is less than prev high, they intersect
		if intervals[i][0] <= compressed[len(compressed)-1][1] {
			// we can merge them into our most recent compressed entry by taking the largest high
			compressed[len(compressed)-1][1] = max(compressed[len(compressed)-1][1], intervals[i][1])
		} else {
			compressed = append(compressed, intervals[i])
		}
	}

	for i := range compressed {
		total += compressed[i][1] - compressed[i][0] + 1
	}

	fmt.Println(total)
}

func parseInput(lines []string) ([][]int, []int) {
	intervals, targets := [][]int{}, []int{}
	i := 0

	for lines[i] != "" {
		str := strings.Split(lines[i], "-")
		low, _ := strconv.Atoi(str[0])
		high, _ := strconv.Atoi(str[1])

		intervals = append(intervals, []int{low, high})
		i++
	}

	i++
	for i < len(lines) {
		target, _ := strconv.Atoi(lines[i])
		targets = append(targets, target)

		i++
	}
	return intervals, targets
}
