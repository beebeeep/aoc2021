package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func showHoles(holes map[point]bool) {
	var w, h int
	for p := range holes {
		if p.x > w {
			w = p.x
		}
		if p.y > h {
			h = p.y
		}
	}
	for x := 0; x <= w; x++ {
		for y := 0; y <= h; y++ {
			p := point{x: x, y: y}
			if holes[p] {
				fmt.Printf("x")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
func fold(holes map[point]bool, folds []point) {
	if len(folds) == 0 {
		return
	}
	f := folds[0]
	for hole := range holes {
		var mirrored point
		if f.x != 0 {
			// fold vertically
			if hole.x < f.x {
				// point is to the left of fold
				continue
			}
			mirrored.y = hole.y
			mirrored.x = 2*f.x - hole.x
		} else {
			// fold horizontally
			if hole.y < f.y {
				// point is above of fold
				continue
			}
			mirrored.x = hole.x
			mirrored.y = 2*f.y - hole.y
		}
		holes[mirrored] = true
		delete(holes, hole)
	}
	fold(holes, folds[1:])
}

func day13() int {
	var (
		holes   = make(map[point]bool)
		folds   = make([]point, 0)
		scanner *bufio.Scanner
	)
	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		tokens := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatal("parsing input: ", err)
		}
		y, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal("parsing input: ", err)
		}
		holes[point{x: x, y: y}] = true
	}
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "=")
		v, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal("parsing input: ", err)
		}
		if strings.HasSuffix(tokens[0], "x") {
			folds = append(folds, point{x: v})
		} else {
			folds = append(folds, point{y: v})
		}
	}
	fold(holes, folds)
	showHoles(holes)
	return len(holes)
}
