package solutions2025

import (
	"fmt"
	"strconv"
)

func Day1Part1(lines []string) {
	curr := 50
	password := 0

	for _, line := range lines {
		direction := line[0]
		magnitude, _ := strconv.Atoi(line[1:])

		if direction == 'L' {
			curr = (curr - magnitude)
			curr = mod(curr, 100)
		} else {
			curr = (curr + magnitude) % 100
		}

		if curr == 0 {
			password++
		}
	}
	fmt.Println(password)
}

func Day1Part2(lines []string) {
	curr := 50
	password := 0

	for _, line := range lines {
		direction := line[0]
		magnitude, _ := strconv.Atoi(line[1:])
		password += magnitude / 100
		magnitude = magnitude % 100

		if direction == 'L' {
			if curr > 0 && curr <= magnitude {
				password++
			}
			curr -= magnitude
		} else {
			curr += magnitude
			if curr >= 100 {
				password++
			}
		}
		curr = mod(curr, 100)
	}

	fmt.Println(password)
}

// positive number only modulo: -1 wraps to 99
func mod(a, b int) int {
	return (a%b + b) % b
}
