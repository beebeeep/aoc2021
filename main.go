package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
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
	fmt.Println(increased)
}

func sum3(a [3]int) int {
	return a[0] + a[1] + a[2]
}
