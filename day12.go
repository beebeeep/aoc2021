package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type cave struct {
	connections []string
	big         bool
}

/*
func countPaths(caves map[string]*cave, from, to string) (numPaths int) {
	if from == to {
		return 1
	}
	caves[from].visited = true
	for _, connection := range caves[from].connections {
		if !caves[connection].visited || caves[connection].big {
			fmt.Printf("going %s -> %s\n", from, connection)
			numPaths += countPaths(caves, connection, to)
		}
	}
	return numPaths
}
*/

func canVisit(caves map[string]*cave, cave string, path []string) bool {
	if caves[cave].big {
		return true
	}
	smallCaveVisits := make(map[string]int)
	for _, c := range path {
		if caves[c].big {
			continue
		}
		smallCaveVisits[c]++
	}
	countTwice := 0
	for _, v := range smallCaveVisits {
		if v > 2 {
			return false
		}
		if v == 2 {
			countTwice++
		}
	}
	return countTwice <= 1
}

func pathContains(path []string, cave string) bool {
	for i := range path {
		if path[i] == cave {
			return true
		}
	}
	return false
}

func findPaths(caves map[string]*cave, from, to string, path []string) [][]string {
	if from == to {
		r := make([]string, len(path)+1)
		copy(r, path)
		r = append(r, to)
		return [][]string{r}
	}
	path = append(path, from)
	newPaths := make([][]string, 0)
	for _, connection := range caves[from].connections {
		if connection == "start" || !canVisit(caves, connection, path) {
			continue
		}

		possiblePaths := findPaths(caves, connection, to, path)
		newPaths = append(newPaths, possiblePaths...)
	}
	return newPaths
}

func addCave(caves map[string]*cave, name string) {
	caves[name] = &cave{
		connections: make([]string, 0, 1),
		big:         name[0] >= 'A' && name[0] <= 'Z',
	}
}

func day12() int {
	var (
		scanner = bufio.NewScanner(os.Stdin)
		caves   = make(map[string]*cave)
	)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "-")
		if len(tokens) != 2 {
			log.Fatal("invalid input")
		}
		name := tokens[0]
		connection := tokens[1]
		if _, ok := caves[name]; !ok {
			addCave(caves, name)
		}
		if _, ok := caves[connection]; !ok {
			addCave(caves, connection)
		}
		caves[name].connections = append(caves[name].connections, connection)
		caves[connection].connections = append(caves[connection].connections, name)
	}
	possiblePaths := findPaths(caves, "start", "end", nil)
	//fmt.Println(possiblePaths)
	return len(possiblePaths)
}
