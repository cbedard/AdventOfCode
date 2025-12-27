package solutions2025

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day9Part1(lines []string) {
	points := linesToPoints(lines)

	area := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			area = max(area, boundedRectArea(points[i], points[j]))
		}
	}
	fmt.Println(area)
}

// Coordinate compression + summed area table. Had to look at reddit for the SAT idea,
func Day9Part2(lines []string) {
	points := linesToPoints(lines)
	grid := NewCompressedGrid(points)

	nCols := grid.X.NumCells()
	nRows := grid.Y.NumCells()

	// hBlocked[i][j] = true if horizontal edge between cell (i,j) and (i,j+1) is blocked
	// vBlocked[i][j] = true if vertical edge between cell (i,j) and (i+1,j) is blocked
	hBlocked := make([][]bool, nCols)
	vBlocked := make([][]bool, nCols-1)
	for i := range nCols {
		hBlocked[i] = make([]bool, nRows-1)
	}
	for i := 0; i < nCols-1; i++ {
		vBlocked[i] = make([]bool, nRows)
	}

	// Mark edges for each boundary segment (including wrap-around)
	for i := range points {
		p1, p2 := points[i], points[(i+1)%len(points)]
		i1, j1 := grid.PointToCell(p1)
		i2, j2 := grid.PointToCell(p2)

		if j1 == j2 { // horizontal segment
			minI, maxI := min(i1, i2), max(i1, i2)
			for ci := minI; ci < maxI; ci++ {
				if j1 > 0 && j1-1 < nRows-1 {
					hBlocked[ci][j1-1] = true
				}
			}
		} else { // vertical segment
			minJ, maxJ := min(j1, j2), max(j1, j2)
			for cj := minJ; cj < maxJ; cj++ {
				if i1 > 0 && i1-1 < nCols-1 {
					vBlocked[i1-1][cj] = true
				}
			}
		}
	}

	// Flood fill from corner (0,0) to mark outside cells
	outside := make([][]bool, nCols)
	for i := range outside {
		outside[i] = make([]bool, nRows)
	}

	type cell struct{ i, j int }
	queue := []cell{{0, 0}}
	outside[0][0] = true

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		if c.j+1 < nRows && !outside[c.i][c.j+1] && !hBlocked[c.i][c.j] {
			outside[c.i][c.j+1] = true
			queue = append(queue, cell{c.i, c.j + 1})
		}
		if c.j > 0 && !outside[c.i][c.j-1] && !hBlocked[c.i][c.j-1] {
			outside[c.i][c.j-1] = true
			queue = append(queue, cell{c.i, c.j - 1})
		}
		if c.i+1 < nCols && !outside[c.i+1][c.j] && !vBlocked[c.i][c.j] {
			outside[c.i+1][c.j] = true
			queue = append(queue, cell{c.i + 1, c.j})
		}
		if c.i > 0 && !outside[c.i-1][c.j] && !vBlocked[c.i-1][c.j] {
			outside[c.i-1][c.j] = true
			queue = append(queue, cell{c.i - 1, c.j})
		}
	}

	// Build summed area table: sat[i][j] = count of outside cells in [0,i) x [0,j)
	sat := make([][]int, nCols+1)
	for i := range sat {
		sat[i] = make([]int, nRows+1)
	}
	for i := range nCols {
		for j := range nRows {
			val := 0
			if outside[i][j] {
				val = 1
			}
			sat[i+1][j+1] = sat[i][j+1] + sat[i+1][j] - sat[i][j] + val
		}
	}

	// Query all pairs of red points
	maxArea := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			i1, j1 := grid.PointToCell(p1)
			i2, j2 := grid.PointToCell(p2)

			// Normalize min/max x/y values
			if i1 > i2 {
				i1, i2 = i2, i1
			}
			if j1 > j2 {
				j1, j2 = j2, j1
			}

			// SAT query: count outside cells in rectangle [i1, i2) x [j1, j2)
			outsideCount := sat[i2][j2] - sat[i1][j2] - sat[i2][j1] + sat[i1][j1]

			if outsideCount == 0 {
				// Valid rectangle - compute area in original coordinates
				area := boundedRectArea(p1, p2)
				maxArea = max(maxArea, area)
			}
		}
	}

	fmt.Println(maxArea)
}

func boundedRectArea(p1, p2 Point) int {
	return (abs(p2.y-p1.y) + 1) * (abs(p2.x-p1.x) + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func linesToPoints(lines []string) []Point {
	points := []Point{}
	for i := range lines {
		str := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(str[0])
		y, _ := strconv.Atoi(str[1])
		points = append(points, Point{y, x})
	}
	return points
}

type CoordCompressor struct {
	values []int       // sorted unique values (original coordinates)
	toIdx  map[int]int // original value -> compressed index
}

// NewCoordCompressor creates a compressor from a set of values. Adds padding values for flood fill boundary detection.
func NewCoordCompressor(vals []int) *CoordCompressor {
	// Dedup and Sort
	set := make(map[int]bool)
	for _, v := range vals {
		set[v] = true
	}

	sorted := make([]int, 0, len(set)+2)
	for v := range set {
		sorted = append(sorted, v)
	}
	slices.Sort(sorted)

	// padding for outside detection
	sorted = append([]int{sorted[0] - 1}, sorted...)
	sorted = append(sorted, sorted[len(sorted)-1]+1)

	// reverse map
	toIdx := make(map[int]int, len(sorted))
	for i, v := range sorted {
		toIdx[v] = i
	}

	return &CoordCompressor{values: sorted, toIdx: toIdx}
}

// ToIndex converts an original coordinate to its compressed grid line index.
func (c *CoordCompressor) ToIndex(val int) int {
	return c.toIdx[val]
}

// ToValue converts a compressed grid line index back to the original coordinate.
func (c *CoordCompressor) ToValue(idx int) int {
	return c.values[idx]
}

// NumCells returns the number of cells (regions between grid lines).
func (c *CoordCompressor) NumCells() int {
	return len(c.values) - 1
}

type CompressedGrid struct {
	X *CoordCompressor
	Y *CoordCompressor
}

// NewCompressedGrid creates a 2D compressed grid from a list of points.
func NewCompressedGrid(points []Point) *CompressedGrid {
	xs := make([]int, len(points))
	ys := make([]int, len(points))
	for i, p := range points {
		xs[i] = p.x
		ys[i] = p.y
	}
	return &CompressedGrid{
		X: NewCoordCompressor(xs),
		Y: NewCoordCompressor(ys),
	}
}

// PointToCell converts an original point to its compressed cell indices.
func (g *CompressedGrid) PointToCell(p Point) (i, j int) {
	return g.X.ToIndex(p.x), g.Y.ToIndex(p.y)
}

// CellToPoint converts compressed grid line indices back to original coordinates.
func (g *CompressedGrid) CellToPoint(i, j int) Point {
	return Point{y: g.Y.ToValue(j), x: g.X.ToValue(i)}
}
