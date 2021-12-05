package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func day5() int {
	var (
		scanner    = bufio.NewScanner(os.Stdin)
		lines      [][2]point
		maxX, maxY int
	)
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		if len(points) != 2 {
			log.Fatal("invalud input")
		}
		line, ok := getLine(points)
		if !ok {
			continue
		}
		lines = append(lines, line)
		if line[0].x > maxX {
			maxX = line[0].x
		}
		if line[1].x > maxX {
			maxX = line[1].x
		}
		if line[0].y > maxY {
			maxY = line[0].y
		}
		if line[1].y > maxY {
			maxY = line[1].y
		}
	}

	field := make([][]int, maxX+1)
	for x := range field {
		field[x] = make([]int, maxY+1)
	}
	for _, line := range lines {
		for _, point := range getLinePoints(line) {
			field[point.x][point.y] += 1
		}
	}
	count := 0
	for x := range field {
		for y := range field[x] {
			if field[x][y] > 1 {
				count++
			}
		}
	}

	return count
}

func getLine(in []string) ([2]point, bool) {
	var (
		r   [2]point
		err error
	)
	for i := 0; i < 2; i++ {
		t := strings.Split(in[i], ",")
		if len(t) != 2 {
			log.Fatal("invalid input")
		}
		r[i].x, err = strconv.Atoi(t[0])
		if err != nil {
			log.Fatal("invalid input")
		}
		r[i].y, err = strconv.Atoi(t[1])
		if err != nil {
			log.Fatal("invalid input")
		}
	}
	// ok := (r[0].x == r[1].x) || (r[0].y == r[1].y)
	ok := (r[0].x == r[1].x) || (r[0].y == r[1].y) || (abs(r[0].x-r[1].x) == abs(r[0].y-r[1].y))

	return r, ok
}

func getLinePoints(line [2]point) []point {
	l := abs(line[0].x - line[1].x)
	if l == 0 {
		l = abs(line[0].y - line[1].y)
	}
	dx := (line[1].x - line[0].x) / l
	dy := (line[1].y - line[0].y) / l
	r := []point{{line[0].x, line[0].y}}
	for i := 1; i <= l; i++ {
		r = append(r, point{r[i-1].x + dx, r[i-1].y + dy})
	}
	return r
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
