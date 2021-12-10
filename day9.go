package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type basin []point

func (b basin) contains(x, y int) bool {
	for _, point := range b {
		if point.x == x && point.y == y {
			return true
		}
	}
	return false
}

func (b *basin) extendAtPoint(heights [][]int, p point) {
	// assuming p already belongs b
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if abs(dx) == 1 && abs(dy) == 1 {
				continue
			}
			x := p.x + dx
			y := p.y + dy
			if h := heights[y][x]; h != 9 && h > p.value && !b.contains(x, y) {
				p1 := point{x: x, y: y, value: h}
				*b = append(*b, p1)
				b.extendAtPoint(heights, p1)
			}
		}
	}
}

func newBasin(heights [][]int, p point) *basin {
	b := &basin{p}
	b.extendAtPoint(heights, p)
	return b
}

func calculateRisk(p []point) int {
	risk := 0
	for _, p := range p {
		risk += 1 + p.value
	}
	return risk
}

func calculateSizes(basins []*basin) int {
	r := 1
	sort.Slice(basins, func(i, j int) bool { return len(*basins[i]) >= len(*basins[j]) })
	for _, b := range basins[:3] {
		r *= len(*b)
	}
	return r
}

func showBasins(heights [][]int, basins []*basin) {
	for y := 1; y < len(heights)-1; y++ {
		for x := 1; x < len(heights[y])-1; x++ {
			basinN := -1
			h := heights[y][x]
			for i, b := range basins {
				if b.contains(x, y) {
					basinN = i
					break
				}
			}
			if basinN != -1 {
				fmt.Printf("\033[1;%dm%d\033[0m", 31+basinN%7, h)
			} else {
				fmt.Printf("%d", h)
			}
		}
		fmt.Println()
	}
}

func day9() int {
	var (
		heights      = make([][]int, 0)
		scanner      = bufio.NewScanner(os.Stdin)
		lowestPoints []point
	)

	for scanner.Scan() {
		rowS := scanner.Text()
		row := make([]int, len(rowS))
		for i, c := range rowS {
			row[i] = int(c - '0')
		}
		heights = append(heights, row)
	}
	for y := 1; y < len(heights)-1; y++ {
		for x := 1; x < len(heights[y])-1; x++ {
			lowest := true
			h := heights[y][x]
			if heights[y-1][x] <= h {
				lowest = false
			}
			if heights[y+1][x] <= h {
				lowest = false
			}
			if heights[y][x-1] <= h {
				lowest = false
			}
			if heights[y][x+1] <= h {
				lowest = false
			}
			if lowest {
				fmt.Printf("\033[1m%d\033[0m", h)
				lowestPoints = append(lowestPoints, point{x: x, y: y, value: h})
			} else {
				fmt.Printf("%d", h)
			}
		}
		fmt.Println()
	}
	fmt.Println("\n\n")

	basins := make([]*basin, 0, len(lowestPoints))
	for _, p := range lowestPoints {
		basins = append(basins, newBasin(heights, p))
	}

	showBasins(heights, basins)
	//return calculateRisk(lowestPoints)
	return calculateSizes(basins)
}
