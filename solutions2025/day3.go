package solutions2025

import (
	"fmt"
	"strconv"
)

// key idea here is that you prioritize
// - taking any digit
// - taking the (leftmost) highest number remaining
// in that order

func Day3Part1(lines []string) {
	total := 0

	for _, line := range lines {
		bestVoltage := 0
		high := byte('0')
		highIndex := 0

		for j := len(line) - 2; j >= 0; j-- {
			if line[j] >= high {
				high = line[j]
				highIndex = j
			}
		}

		for j := len(line) - 1; j > highIndex; j-- {
			curr, _ := strconv.Atoi(string(high) + string(line[j]))
			bestVoltage = max(bestVoltage, curr)
		}

		total += bestVoltage
	}

	fmt.Println(total)
}

func Day3Part2(lines []string) {
	total := 0

	for _, line := range lines {
		voltage, _ := strconv.Atoi(lengthNVoltage(12, line, ""))
		total += voltage
	}

	fmt.Println(total)
}

func lengthNVoltage(N int, battery, curr string) string {
	if len(curr) == N {
		return curr
	}

	high := byte('0')
	highIndex := 0

	for i := 0; i < len(battery)-(N-len(curr)-1); i++ {
		if battery[i] > high {
			high = battery[i]
			highIndex = i
		}
	}

	curr += string(battery[highIndex])
	return lengthNVoltage(N, battery[highIndex+1:], curr)
}
