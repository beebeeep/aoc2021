package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ебитес(рыбы *[]uint8) {
	žuvys := *рыбы
	kiekŽuvų := len(žuvys)
	for i := 0; i < kiekŽuvų; i++ {
		switch žuvys[i] {
		case 0:
			žuvys[i] = 6
			*рыбы = append(*рыбы, 8)
			žuvys = *рыбы
		default:
			žuvys[i]--
		}
	}
}

func day6() int {
	var (
		scanner = bufio.NewScanner(os.Stdin)
		рыбы    = make([]uint8, 0, 1000)
	)
	scanner.Scan()
	for _, t := range strings.Split(scanner.Text(), ",") {
		n, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal("invalid input: ", err)
		}
		рыбы = append(рыбы, uint8(n))
	}

	for day := 0; day < 256; day++ {
		fmt.Printf("денб %d, рыбов %d\n", day, len(рыбы))
		ебитес(&рыбы)
	}

	return len(рыбы)
}
