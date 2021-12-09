package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func day1() int {
	scanner := bufio.NewScanner(os.Stdin)
	var (
		i         = -1
		window    [3]int
		lastSum   int
		increased int
	)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("incorrect input", err)
		}
		i++
		window[i%3] = num
		if i < 3 {
			continue
		}
		sum := sum3(window)
		if sum > lastSum {
			increased++
		}
		lastSum = sum
	}

	return increased
}
