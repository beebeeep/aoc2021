package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var BIT_SIZE = 12

func day3() int {
	var (
		scanner     = bufio.NewScanner(os.Stdin)
		values      []int
		oxygen, co2 int
	)
	for scanner.Scan() {
		input := scanner.Text()
		value, err := strconv.ParseInt(input, 2, 64)
		if err != nil {
			log.Fatal("cannot parse number: ", err)
		}
		values = append(values, int(value))

	}
	oxygen = findByCriteria(values, true, 0)
	co2 = findByCriteria(values, false, 0)
	return oxygen * co2
}

func findByCriteria(values []int, isOxy bool, pos int) int {
	if len(values) == 1 {
		return values[0]
	}
	if len(values) == 0 || pos > BIT_SIZE {
		log.Fatal("no match?")
		return 0
	}

	var (
		sum        int
		wantOnes   bool
		half       = (len(values) + 1) / 2
		nextValues = make([]int, 0, len(values)/2)
	)
	for _, v := range values {
		if v&(1<<(BIT_SIZE-1-pos)) != 0 {
			sum++
		}
	}
	if sum >= half {
		// more or equally ones
		wantOnes = isOxy
	} else if sum < half {
		// more zeros
		wantOnes = !isOxy
	}

	for _, v := range values {
		if v&(1<<(BIT_SIZE-1-pos)) != 0 {
			if wantOnes {
				nextValues = append(nextValues, v)
			}
		} else {
			if !wantOnes {
				nextValues = append(nextValues, v)
			}
		}
	}
	return findByCriteria(nextValues, isOxy, pos+1)
}
