package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func day2() int {
	var (
		scanner         = bufio.NewScanner(os.Stdin)
		posX, posZ, aim int
	)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if len(tokens) != 2 {
			log.Fatal("invalid input:", tokens)
		}
		value, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal("cannot parse number:", err)
		}
		switch tokens[0] {
		case "down":
			aim += value
		case "up":
			aim -= value
		case "forward":
			posX += value
			posZ += aim * value
		}
	}
	return posX * posZ
}
