package solutions2025

import (
	"fmt"
	"strings"
)

type Point struct {
	y, x int
}

func Day7Part1(grid []string) {
	// general pattern -- we keep a list of current beams in a map[Point]bool
	// iterate across all beams, try to move one downward. if it hits a beam we split it
	// all beams move 1 down each turn, we can put this inside of a loop len(grid)
	sIndex := strings.Index(grid[0], "S")
	beams := map[Point]bool{Point{0, sIndex}: true}
	splitCount := 0

	for i := 0; i < len(grid)-1; i++ {
		newBeams := make(map[Point]bool)
		for beam := range beams {
			if grid[beam.y+1][beam.x] == '^' {
				newBeams[Point{beam.y + 1, beam.x - 1}] = true
				newBeams[Point{beam.y + 1, beam.x + 1}] = true
				splitCount++
			} else {
				newBeams[Point{beam.y + 1, beam.x}] = true
			}
		}
		beams = newBeams
	}

	fmt.Println(splitCount)
}

func Day7Part2(grid []string) {
	// if theres 35 beams on a given point y,x -- we either add 35 points to that y+1,x or 35 to the split points -> sum the bottom row

	sIndex := strings.Index(grid[0], "S")
	beams := map[Point]bool{Point{0, sIndex}: true}

	beamCache := make([][]int, len(grid))
	for i := range beamCache {
		beamCache[i] = make([]int, len(grid[i]))
	}
	beamCache[0][sIndex] = 1

	for i := 0; i < len(grid)-1; i++ {
		newBeams := make(map[Point]bool)
		for beam := range beams {
			if grid[beam.y+1][beam.x] == '^' {
				newBeams[Point{beam.y + 1, beam.x - 1}] = true
				newBeams[Point{beam.y + 1, beam.x + 1}] = true

				beamCache[beam.y+1][beam.x-1] += beamCache[beam.y][beam.x]
				beamCache[beam.y+1][beam.x+1] += beamCache[beam.y][beam.x]
			} else {
				newBeams[Point{beam.y + 1, beam.x}] = true
				beamCache[beam.y+1][beam.x] += beamCache[beam.y][beam.x]
			}
		}
		beams = newBeams
	}

	total := 0
	for i := range beamCache[len(beamCache)-1] {
		total += beamCache[len(beamCache)-1][i]
	}

	fmt.Println(total)
}
