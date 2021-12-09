package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func guessPattern(patterns []string) map[string]int {
	var (
		r          = make(map[string]int)
		result     = make(map[string]int)
		p1, p4, p7 map[rune]bool
	)
	for _, pattern := range patterns {
		switch len(pattern) {
		case 2:
			r[pattern] = 1
			p1 = str2map(pattern)
		case 4:
			r[pattern] = 4
			p4 = str2map(pattern)
		case 3:
			r[pattern] = 7
			p7 = str2map(pattern)
		case 7:
			r[pattern] = 8
		}
	}

	for _, pattern := range patterns {
		switch len(pattern) {
		case 6:
			r[pattern] = check069(pattern, p1, p4, p7)
		case 5:
			r[pattern] = check235(pattern, p1, p4, p7)
		}
	}
	for k, v := range r {
		result[sortLetters(k)] = v
	}

	return result
}

func check069(s string, p1, p4, p7 map[rune]bool) int {
	is0 := true
	is9 := true
	is6 := false

	smap := str2map(s)
	for c := range p7 {
		is0 = is0 && smap[c]
		is9 = is9 && smap[c]
	}
	for c := range p4 {
		is9 = is9 && smap[c]
		is6 = is6 || !smap[c]
	}
	if is9 {
		return 9
	}
	if is0 {
		return 0
	}
	if is6 {
		return 6
	}
	log.Fatal("wot?")
	return 0
}

func check235(s string, p1, p4, p7 map[rune]bool) int {
	var (
		is5    = true
		is3    = true
		runes5 = make(map[rune]bool)
		smap   = str2map(s)
	)
	for c := range p4 {
		if !p1[c] {
			runes5[c] = true
		}
	}

	for c := range runes5 {
		is5 = is5 && smap[c]
	}
	if is5 {
		return 5
	}
	for c := range p1 {
		is3 = is3 && smap[c]
	}
	if is3 {
		return 3
	}

	return 2
}

func decodeOutput(value []string, mapping map[string]int) (r int) {
	r += mapping[sortLetters(value[0])] * 1000
	r += mapping[sortLetters(value[1])] * 100
	r += mapping[sortLetters(value[2])] * 10
	r += mapping[sortLetters(value[3])]
	return r
}

func count1478(values [][]string) int {
	var count int
	for _, value := range values {
		for _, digit := range value {
			if len(digit) == 2 || len(digit) == 4 || len(digit) == 3 || len(digit) == 7 {
				count++
			}
		}
	}
	return count
}

func day8() int {
	var (
		scanner  = bufio.NewScanner(os.Stdin)
		patterns = make([][]string, 0, 200)
		output   = make([][]string, 0, 200)
	)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " | ")
		patterns = append(patterns, strings.Split(input[0], " "))
		output = append(output, strings.Split(input[1], " "))
	}
	sum := 0
	for i, pattern := range patterns {
		sum += decodeOutput(output[i], guessPattern(pattern))
	}
	return sum
}
