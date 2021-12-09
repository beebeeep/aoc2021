package main

import "sort"

type point struct {
	x, y, value int
}

func sortLetters(s string) string {
	a := []rune(s)
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	return string(a)
}

func str2map(s string) map[rune]bool {
	r := make(map[rune]bool)
	for _, c := range s {
		r[c] = true
	}
	return r
}

func sum3(a [3]int) int {
	return a[0] + a[1] + a[2]
}
