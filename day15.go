package main

import (
	"bufio"
	"fmt"
	"os"
)

type rambler struct {
	path     []point
	position point
}

func wasHere(path []point, x, y int) bool {
	for _, p := range path {
		if p.x == x && p.y == y {
			return true
		}
	}
	return false
}
func showPath(cave [][]int, path []point) {
	for y := range cave {
		for x := range cave[y] {
			if wasHere(path, x, y) {
				fmt.Printf("\033[1m%d\033[0m", cave[y][x])
			} else {
				fmt.Printf("%d", cave[y][x])
			}
		}
		fmt.Printf("\n")
	}
}

func getPathCost(p []point) int {
	cost := 0
	for _, c := range p[1:] {
		cost += c.value
	}
	return cost
}

func day15() int {
	cave := make([][]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		r := make([]int, 0)
		for _, c := range scanner.Text() {
			r = append(r, int(c-'0'))
		}
		cave = append(cave, r)
	}

	startPoint := point{x: 0, y: 0, value: cave[0][0]}
	visitedPoints := map[point]bool{
		startPoint: true,
	}

	ramblers := map[string]*rambler{
		"rambler-0": {path: []point{startPoint}, position: startPoint},
	}
	step := 0
	ramblerCount := 0
	paths := make([][]point, 0)
	endX := len(cave[0]) - 1
	endY := len(cave) - 1
RAMBLE:
	for {
		step++
		if len(ramblers) == 0 {
			break RAMBLE
		}
		for name, r := range ramblers {
			if r.position.x == endX && r.position.y == endY {
				paths = append(paths, r.path)
				delete(ramblers, name)
			}
			r.position.value--
			if r.position.value > 0 {
				continue
			}
			for _, delta := range [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
				x := r.position.x + delta[0]
				y := r.position.y + delta[1]
				if x < 0 || y < 0 || x > endX || y > endY {
					continue
				}
				p := point{x: x, y: y, value: cave[y][x]}
				if visitedPoints[p] {
					continue
				}
				visitedPoints[p] = true
				newRambler := &rambler{
					path:     make([]point, len(r.path)+1),
					position: p,
				}
				copy(newRambler.path, r.path)
				newRambler.path = append(newRambler.path, p)

				ramblerCount++
				ramblers[fmt.Sprintf("rambler-%d", ramblerCount)] = newRambler
			}
			delete(ramblers, name)
		}
	}

	minCost := 0

	for i := range paths {
		/*
			showPath(cave, paths[i])
			fmt.Println()
		*/
		if cost := getPathCost(paths[i]); minCost == 0 || cost < minCost {
			minCost = cost
		}
	}
	return minCost

}
