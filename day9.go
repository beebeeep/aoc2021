package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	fmt.Println(len(heights), len(heights[0]))
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
	risk := 0
	for _, p := range lowestPoints {
		risk += 1 + p.value
	}
	return risk

}
