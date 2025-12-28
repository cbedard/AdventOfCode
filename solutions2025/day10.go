package solutions2025

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Day10Part1(lines []string) {
	sum := 0
	for _, line := range lines {
		lights, wirings, _ := parseMachine(line)

		// idea here is brute force search (DFS) + cycle detection, lights <= 10
		// cycle detection can be implmeneted with map[lights]presses, we only stop if we hit a cycle at >= pre-existing presses
		cycles := make(map[string]int)
		curr := ""
		for range len(lights) {
			curr += "."
		}

		DFSLights(lights, curr, wirings, 0, cycles)
		sum += cycles[lights]
	}

	fmt.Println(sum)
}

func DFSLights(target, curr string, wirings [][]int, presses int, cycles map[string]int) {
	prev, ok := cycles[curr]
	if ok && prev <= presses {
		return
	}
	cycles[curr] = presses

	for i := range wirings {
		next := []byte(curr)

		for _, light := range wirings[i] {
			if next[light] == '.' {
				next[light] = '#'
			} else {
				next[light] = '.'
			}
		}

		DFSLights(target, string(next), wirings, presses+1, cycles)
	}
}

// --- Part 2 ---
// At each step: solve for parities (gaussian elim), subtract, divide by 2, recurse
func Day10Part2(lines []string) {
	sum := 0

	for _, line := range lines {
		_, wirings, joltages := parseMachine(line)

		presses := solveByHalving(wirings, joltages)
		sum += presses
	}

	fmt.Println(sum)
}

// solveByHalving uses binary approach: solve LSB, subtract, divide by 2, recurse
func solveByHalving(wirings [][]int, target []int) int {
	numCounters := len(target)
	numButtons := len(wirings)

	// buttonMatrix[i][j] = 1 if button j affects counter i
	buttonMatrix := make([][]int, numCounters)
	for i := range numCounters {
		buttonMatrix[i] = make([]int, numButtons)
		for j, button := range wirings {
			for _, counter := range button {
				if counter == i {
					buttonMatrix[i][j] = 1
				}
			}
		}
	}

	return solveRecursive(wirings, buttonMatrix, target, 1)
}

// solveRecursive tries all valid GF(2) solutions at current level and recurses
func solveRecursive(wirings [][]int, buttonMatrix [][]int, current []int, multiplier int) int {
	if allZero(current) {
		return 0
	}

	numCounters, numButtons := len(current), len(wirings)

	parities := make([]int, numCounters)
	for i, v := range current {
		parities[i] = v % 2
	}

	// Get all valid GF(2) solutions
	solutions := solveGF2All(buttonMatrix, parities, numButtons)
	bestTotal := 1 << 30

	for _, presses := range solutions {
		next := make([]int, numCounters)
		copy(next, current)
		levelCost := 0

		for j, p := range presses {
			levelCost += p
			if p == 1 {
				for _, counter := range wirings[j] {
					next[counter] -= 1
				}
			}
		}

		// negative values -> invalid path
		if slices.ContainsFunc(next, func(v int) bool {
			return v < 0
		}) {
			continue
		}

		for i := range next {
			next[i] /= 2
		}

		futureCost := solveRecursive(wirings, buttonMatrix, next, multiplier*2)
		totalCost := levelCost*multiplier + futureCost

		if totalCost < bestTotal {
			bestTotal = totalCost
		}
	}

	return bestTotal
}

// solveGF2All returns all valid solutions to Ax = b over GF(2)
func solveGF2All(matrix [][]int, target []int, numButtons int) [][]int {
	numRows := len(matrix)

	// Create augmented matrix [A|b]
	aug := make([][]int, numRows)
	for i := range numRows {
		aug[i] = make([]int, numButtons+1)
		copy(aug[i], matrix[i])
		aug[i][numButtons] = target[i]
	}

	// Gaussian elimination over GF(2)
	pivotCol, pivotRow := 0, 0
	pivotCols := make([]int, 0)
	freeCols := make([]int, 0)

	for pivotCol < numButtons && pivotRow < numRows {
		found := -1
		for i := pivotRow; i < numRows; i++ {
			if aug[i][pivotCol] == 1 {
				found = i
				break
			}
		}

		if found == -1 {
			freeCols = append(freeCols, pivotCol)
			pivotCol++
			continue
		}

		aug[pivotRow], aug[found] = aug[found], aug[pivotRow]
		pivotCols = append(pivotCols, pivotCol)

		for i := range numRows {
			if i != pivotRow && aug[i][pivotCol] == 1 {
				for j := 0; j <= numButtons; j++ {
					aug[i][j] ^= aug[pivotRow][j]
				}
			}
		}

		pivotRow++
		pivotCol++
	}

	for c := pivotCol; c < numButtons; c++ {
		freeCols = append(freeCols, c)
	}

	// Check for row with all zero coefficients but non-zero target
	for i := pivotRow; i < numRows; i++ {
		if aug[i][numButtons] == 1 && allZero(aug[i][:numButtons]) {
			return nil // inconsistent
		}
	}

	// Generate all solutions by trying all free variable combinations
	numFree := len(freeCols)
	var solutions [][]int

	for mask := 0; mask < (1 << numFree); mask++ {
		solution := make([]int, numButtons)

		for i, col := range freeCols {
			if (mask>>i)&1 == 1 {
				solution[col] = 1
			}
		}

		for i := len(pivotCols) - 1; i >= 0; i-- {
			row := i
			col := pivotCols[i]

			val := aug[row][numButtons]
			for j := col + 1; j < numButtons; j++ {
				val ^= (solution[j] * aug[row][j])
			}
			solution[col] = val
		}

		solutions = append(solutions, solution)
	}

	return solutions
}

// --- Utils ---

func allZero(arr []int) bool {
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

func parseMachine(line string) (string, [][]int, []int) {
	// Parse indicator light diagram [...]
	bracketRe := regexp.MustCompile(`\[([^\]]+)\]`)
	bracketMatch := bracketRe.FindStringSubmatch(line)
	var indicator string
	if len(bracketMatch) > 1 {
		indicator = bracketMatch[1]
	}

	// Parse button wiring schematics (...)
	parenRe := regexp.MustCompile(`\(([^)]+)\)`)
	parenMatches := parenRe.FindAllStringSubmatch(line, -1)
	var wirings [][]int
	for _, match := range parenMatches {
		if len(match) > 1 {
			nums := parseIntList(match[1])
			wirings = append(wirings, nums)
		}
	}

	// Parse joltage requirements {...}
	braceRe := regexp.MustCompile(`\{([^}]+)\}`)
	braceMatch := braceRe.FindStringSubmatch(line)
	var joltages []int
	if len(braceMatch) > 1 {
		joltages = parseIntList(braceMatch[1])
	}

	return indicator, wirings, joltages
}

func parseIntList(s string) []int {
	var result []int
	parts := strings.Split(s, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if n, err := strconv.Atoi(p); err == nil {
			result = append(result, n)
		}
	}
	return result
}
