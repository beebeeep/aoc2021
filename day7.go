package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateFuel(crabs []int, targetPos int) int {
	var fuel int
	for _, pos := range crabs {
		fuel += fuelCost(abs(pos - targetPos))
	}
	return fuel
}

func fuelCost(dist int) int {
	return dist * (1 + dist) / 2
}

func day7() int {
	var (
		scanner             = bufio.NewScanner(os.Stdin)
		crabs               = make([]int, 0, 1000)
		sum, fuel, max, min int
	)
	scanner.Scan()
	for _, t := range strings.Split(scanner.Text(), ",") {
		n, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal("invalid input: ", err)
		}
		crabs = append(crabs, n)
		if n < min || min == 0 {
			min = n
		}
		if n > max {
			max = n
		}
		sum += n
	}
	for i := min; i <= max; i++ {
		if t := calculateFuel(crabs, i); t < fuel || fuel == 0 {
			fuel = t
		}
	}

	return fuel
}
