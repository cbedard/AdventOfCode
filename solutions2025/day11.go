package solutions2025

import (
	"fmt"
	"strings"
)

func Day11Part1(lines []string) {
	adj := parse(lines)
	fmt.Println(countPaths(adj, "you", "out"))
}

// this solution leans on some manual testing with countPaths, theres no paths where dac comes before fft
func Day11Part2(lines []string) {
	adj := parse(lines)
	fmt.Println(countPaths(adj, "svr", "fft") * countPaths(adj, "fft", "dac") * countPaths(adj, "dac", "out"))
}

func countPaths(adj map[string][]string, src, dest string) int {
	memo := make(map[string]int)

	var dfs func(node string) int
	dfs = func(node string) int {
		if node == dest {
			return 1
		}
		if val, found := memo[node]; found {
			return val
		}

		count := 0
		for _, next := range adj[node] {
			count += dfs(next)
		}

		memo[node] = count
		return count
	}

	return dfs(src)
}

// ccc: ddd eee fff
func parse(lines []string) map[string][]string {
	adj := make(map[string][]string)

	for _, line := range lines {
		str := strings.Split(line, ":")
		src := str[0]
		neighbors := strings.Split(str[1], " ")[1:]

		adj[src] = append(adj[src], neighbors...)
	}
	return adj
}
