package solutions2025

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2Part1(lines []string) {
	tokens := strings.Split(lines[0], ",")
	invalid := 0

	for i := range tokens {
		nums := strings.Split(tokens[i], "-")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])

		for n1 <= n2 {
			str1 := strconv.Itoa(n1)
			strLen := len(str1)

			if strLen%2 == 0 && str1[:strLen/2] == str1[strLen/2:] {
				//fmt.Println(str1[:strLen/2], str1[strLen/2:])
				invalid += n1
			}
			n1++
		}
	}

	fmt.Println(invalid)
}

// Part 2 notes
// for there to be a repeatable sequence of length y in str, len(str) % y == 0
func Day2Part2(lines []string) {
	tokens := strings.Split(lines[0], ",")
	invalid := 0

	for i := range tokens {
		nums := strings.Split(tokens[i], "-")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])

		for n1 <= n2 {
			str1 := strconv.Itoa(n1)
			strLen := len(str1)

		pieceSearch:
			for pieces := 2; pieces <= strLen; pieces++ {
				if strLen%pieces != 0 {
					continue
				}

				//copy str, take len pieceSize, compare it to nexrt pieceSize
				newStr := str1
				pieceSize := strLen / pieces

				currPiece := newStr[:pieceSize]
				newStr = newStr[pieceSize:]

				for range pieces - 1 {
					if currPiece != newStr[:pieceSize] {
						continue pieceSearch
					} else {
						newStr = newStr[pieceSize:]
					}
				}

				//if we made it here the id is invalid
				invalid += n1
				break pieceSearch
			}
			n1++
		}
	}

	fmt.Println(invalid)
}
