package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const BOARD_SIZE = 5

type field struct {
	number int
	marked bool
}
type board [BOARD_SIZE][BOARD_SIZE]field

func day4() int {
	doneBoards := 0
	numbers, boards := parseDay4Input()
	for _, n := range numbers {
		mark(boards, n)
		for i, board := range boards {
			if board == nil {
				continue
			}
			if win(board) {
				doneBoards++
				if doneBoards == len(boards) {
					return score(board, n)
				}
				boards[i] = nil
			}
		}
	}
	return 0
}

func score(board *board, n int) int {
	sum := 0
	if board == nil {
		return 0
	}
	for _, row := range board {
		for _, column := range row {
			if !column.marked {
				sum += column.number
			}
		}
	}
	return n * sum
}

func mark(boards []*board, n int) {
	for i := range boards {
		if boards[i] == nil {
			continue
		}
		for r := range boards[i] {
			for c := range boards[i][r] {
				if boards[i][r][c].number == n {
					boards[i][r][c].marked = true
				}
			}
		}
	}
}

func win(board *board) bool {

	for i := 0; i < BOARD_SIZE; i++ {
		// rows
		win := true
		for _, column := range board[i] {
			win = win && column.marked
		}
		if win {
			return true
		}
		// columns
		win = true
		for _, row := range board {
			win = win && row[i].marked
		}
		if win {
			return true
		}
	}
	return false
}

func showBoard(board *board) {
	for _, row := range board {
		fmt.Printf("%2d %2d %2d %2d %2d\n", row[0].number, row[1].number, row[2].number, row[3].number, row[4].number)
	}
}

func parseDay4Input() ([]int, []*board) {
	var (
		boards  []*board
		numbers []int
		scanner = bufio.NewScanner(os.Stdin)
	)

	scanner.Scan()
	for _, t := range strings.Split(scanner.Text(), ",") {
		n, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal("paring input: ", err)
		}
		numbers = append(numbers, n)
	}

	scanner.Scan()
	var row, boardI int
	boards = append(boards, &board{})
	lineRe := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boardI++
			row = 0
			boards = append(boards, &board{})
			continue
		}
		for column, t := range lineRe.FindStringSubmatch(line)[1:] {
			n, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal("paring input: ", err)
			}
			boards[boardI][row][column].number = n
		}
		row++
	}

	return numbers, boards
}
