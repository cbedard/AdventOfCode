package main

import (
	//"AOC/solutions2022"
	//"AOC/solutions2023"

	"AOC/solutions2025"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/home/cameron/Documents/AdventOfCode/input/day5.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	solutions2025.Day5Part2(lines)
}
