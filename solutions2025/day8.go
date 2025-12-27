package solutions2025

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Point3D struct {
	x, y, z int
}

type Edge3D struct {
	node1, node2 int
	dist         float64
}

func Day8Part1(lines []string) {
	nodes, edges := nodesAndEdges(lines)

	dsu := NewDSU(len(nodes))
	for i := range 1000 {
		dsu.Union(edges[i].node1, edges[i].node2)
	}

	groupMap := make(map[int]int)
	for i := range nodes {
		groupMap[dsu.Find(i)]++
	}

	// largest 3 groups
	sizes := make([]int, 0, len(groupMap))
	for _, size := range groupMap {
		sizes = append(sizes, size)
	}
	slices.SortFunc(sizes, func(a, b int) int { return b - a })

	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}

func Day8Part2(lines []string) {
	nodes, edges := nodesAndEdges(lines)

	dsu := NewDSU(len(nodes))

	for i := range edges {
		dsu.Union(edges[i].node1, edges[i].node2)

		if dsu.MaxSize == len(nodes) {
			fmt.Println(nodes[edges[i].node1].x * nodes[edges[i].node2].x)
			break
		}
	}
}

func nodesAndEdges(lines []string) ([]Point3D, []Edge3D) {
	nodes := linesToNodes(lines)
	edges := []Edge3D{}

	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {
			edges = append(edges, Edge3D{i, j, distance3D(nodes[i], nodes[j])})
		}
	}

	slices.SortFunc(edges, func(i, j Edge3D) int {
		if i.dist > j.dist {
			return 1
		} else if j.dist > i.dist {
			return -1
		}
		return 0
	})

	return nodes, edges
}

func linesToNodes(lines []string) []Point3D {
	nodes := []Point3D{}
	for i := range lines {
		str := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(str[0])
		y, _ := strconv.Atoi(str[1])
		z, _ := strconv.Atoi(str[2])

		nodes = append(nodes, Point3D{x, y, z})
	}
	return nodes
}

func distance3D(p1, p2 Point3D) float64 {
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)
	dz := float64(p1.z - p2.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// --- Disjoint Set Union w/ size instead of rank ---
type DSU struct {
	parent  []int
	size    []int
	MaxSize int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	maxSize := 1

	return &DSU{parent: parent, size: size, MaxSize: maxSize}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
	px, py := d.Find(x), d.Find(y)
	if px == py {
		return false
	}

	if d.size[px] < d.size[py] {
		px, py = py, px
	}

	d.parent[py] = px
	d.size[px] += d.size[py]
	d.MaxSize = max(d.MaxSize, d.size[px])

	return true
}
