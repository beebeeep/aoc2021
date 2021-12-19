package main

import (
	"bufio"
	"fmt"
	"os"
)

type rambler struct {
	path     []point
	score    int
	position point
}

func extendCave(cave [][]int) [][]int {
	result := make([][]int, len(cave)*5)
	for y := range cave {
		for dy := 0; dy < 5; dy++ {
			y1 := y + dy*len(cave)
			result[y1] = make([]int, len(cave[y])*5)
			for x := range cave[y] {
				for dx := 0; dx < 5; dx++ {
					x1 := x + dx*len(cave[y])
					result[y1][x1] = cave[y][x] + dx + dy
					if result[y1][x1] > 9 {
						result[y1][x1] = result[y1][x1] % 9
					}

				}
			}
		}
	}
	return result
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

	cave = extendCave(cave)
	startPoint := point{x: 0, y: 0, value: cave[0][0]}
	visitedPoints := map[point]int{
		startPoint: cave[0][0],
	}
	ramblers := map[string]*rambler{
		"rambler-0": {path: []point{startPoint}, position: startPoint},
	}
	ramblerCount := 0
	endX := len(cave[0]) - 1
	endY := len(cave) - 1
	for {
		for name, r := range ramblers {
			if r.position.x == endX && r.position.y == endY {
				return r.score
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
				if v, ok := visitedPoints[p]; ok && v <= r.score+p.value {
					continue
				}
				visitedPoints[p] = r.score + p.value
				newRambler := &rambler{
					path:     make([]point, len(r.path)+1),
					position: p,
					score:    p.value + r.score,
				}
				copy(newRambler.path, r.path)
				newRambler.path = append(newRambler.path, p)

				ramblerCount++
				ramblers[fmt.Sprintf("rambler-%d", ramblerCount)] = newRambler
			}
			delete(ramblers, name)
		}
	}
}
