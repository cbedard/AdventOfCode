package main

import (
	//"AOC/solutions2022"
	//"AOC/solutions2023"

	"AOC/solutions2025"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Open("/home/cameron/Documents/AdventOfCode/input/day10.txt")
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

	start := time.Now()
	solutions2025.Day10Part2(lines)
	fmt.Println("Wall time:", time.Since(start))
}
