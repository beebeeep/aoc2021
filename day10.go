package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type runeStack []rune

func newRuneStack() runeStack {
	return make(runeStack, 0)
}

func (s *runeStack) push(r rune) {
	*s = append(*s, r)
}

func (s *runeStack) pop() (r rune) {
	l := len(*s)
	if l == 0 {
		return 0
	}
	r = (*s)[l-1]
	*s = (*s)[:l-1]
	return r
}

func getPair(r rune) rune {
	switch r {
	case '{':
		return '}'
	case '}':
		return '{'
	case '<':
		return '>'
	case '>':
		return '<'
	case '(':
		return ')'
	case ')':
		return '('
	case '[':
		return ']'
	case ']':
		return '['
	}
	return 0
}

var (
	getCorruptedScore = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	incompleteScore = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

func parse(line string) ([]rune, rune) {
	expr := newRuneStack()
	for _, c := range []rune(line) {
		//fmt.Printf("%c | %s\n", c, string(expr))
		switch c {
		case '[', '<', '{', '(':
			expr.push(c)
		case ']', '>', '}', ')':
			if expr.pop() != getPair(c) {
				return nil, c
			}
		}
	}
	incomplete := newRuneStack()
	for _, c := range expr {
		incomplete.push(getPair(c))
	}
	return incomplete, 0
}

func getIncompleteScore(s runeStack) (r int) {
	// runeStack is in reversed order
	for i := len(s) - 1; i >= 0; i-- {
		r *= 5
		r += incompleteScore[s[i]]
	}
	return r
}

func day10() int {
	var (
		scanner         = bufio.NewScanner(os.Stdin)
		corruptedScore  int
		incompleteScore = make([]int, 0)
	)

	for scanner.Scan() {
		line := scanner.Text()
		incomplete, corrupted := parse(line)
		fmt.Printf("Line %s, incomplete: %s (%d), corrupted: %d\n", line, string(incomplete), getIncompleteScore(incomplete), corrupted)
		corruptedScore += getCorruptedScore[corrupted]
		if len(incomplete) > 0 {
			incompleteScore = append(incompleteScore, getIncompleteScore(incomplete))
		}
	}
	sort.Ints(incompleteScore)
	return incompleteScore[len(incompleteScore)/2]
}
