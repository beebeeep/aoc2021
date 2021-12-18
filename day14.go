package main

import (
	"bufio"
	"os"
	"strings"
)

func unfoldRules(rules map[string]byte) map[string][]string {
	result := make(map[string][]string)
	for pair, c := range rules {
		result[pair] = []string{
			string([]byte{pair[0], c}),
			string([]byte{c, pair[1]}),
		}
	}
	return result
}

func splitToPairs(input string) map[string]uint64 {
	result := make(map[string]uint64)
	for i := 0; i < len(input)-1; i++ {
		result[input[i:i+2]]++
	}
	return result
}

func day14() uint64 {
	var (
		formula  string
		rawRules = make(map[string]byte)
		min      uint64
		max      uint64
	)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	formula = scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " -> ")
		rawRules[tokens[0]] = tokens[1][0]
	}
	rules := unfoldRules(rawRules)
	pairs := splitToPairs(formula)

	elements := make(map[byte]uint64)
	for _, c := range []byte(formula) {
		elements[c]++
	}

	for step := 0; step < 40; step++ {
		newPairs := make(map[string]uint64)
		for pair, v := range pairs {
			rule := rules[pair]
			elements[rule[0][1]] += v
			newPairs[rule[0]] += v
			newPairs[rule[1]] += v
		}
		pairs = newPairs
	}
	for _, v := range elements {
		if min == 0 || v < min {
			min = v
		}
		if max == 0 || v > max {
			max = v
		}
	}
	return (max - min)

}
