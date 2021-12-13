package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type dumbo struct {
	energy  int
	flashed bool
}

func dumboFlash(dumbos [][]dumbo, x0, y0 int) (flashes int) {
	dumbos[y0][x0].flashed = true
	flashes++
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			x := x0 + dx
			y := y0 + dy
			if x < 0 || x >= len(dumbos[y0]) || y < 0 || y >= len(dumbos) || (dx == 0 && dy == 0) {
				continue
			}
			dumbos[y][x].energy++
			if dumbos[y][x].energy > 9 && !dumbos[y][x].flashed {
				flashes += dumboFlash(dumbos, x, y)
			}
		}
	}
	return flashes
}
func showDumbos(dumbos [][]dumbo) {
	for y := range dumbos {
		for x := range dumbos[y] {
			if dumbos[y][x].energy == 0 {
				fmt.Printf("\033[1;39m%d\033[0m", dumbos[y][x].energy)
			} else if dumbos[y][x].energy == 9 {
				fmt.Printf("\033[1;31m%d\033[0m", dumbos[y][x].energy)
			} else {
				fmt.Printf("%d", dumbos[y][x].energy)
			}
		}
		fmt.Println()
	}
}

func dumboStep(dumbos [][]dumbo) (flashes int, allFlashed bool) {
	for y := range dumbos {
		for x := range dumbos[y] {
			dumbos[y][x].energy++
			dumbos[y][x].flashed = false
		}
	}
	for y := range dumbos {
		for x := range dumbos[y] {
			if dumbos[y][x].energy > 9 && !dumbos[y][x].flashed {
				flashes += dumboFlash(dumbos, x, y)
			}
		}
	}
	allFlashed = true
	for y := range dumbos {
		for x := range dumbos[y] {
			if dumbos[y][x].flashed {
				dumbos[y][x].energy = 0
			} else {
				allFlashed = false
			}
		}
	}
	return flashes, allFlashed
}

func day11() int {
	var (
		scanner = bufio.NewScanner(os.Stdin)
		dumbos  [][]dumbo
		flashes int
	)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]dumbo, len(line))
		for i, d := range line {
			row[i].energy = int(d - '0')
		}
		dumbos = append(dumbos, row)
	}
	for i := 0; i < 1000; i++ {
		fmt.Printf("Step %d\n", i)
		showDumbos(dumbos)
		n, allFlashed := dumboStep(dumbos)
		flashes += n
		if allFlashed {
			return i + 1
		}
		time.Sleep(100 * time.Millisecond)

	}
	return flashes

}
