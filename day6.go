package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ебитес(рыбы *[9]uint64) {
	žuvys := *рыбы
	for день := 1; день <= 8; день++ {
		žuvys[день-1] = (*рыбы)[день]
		if день < 8 {
			žuvys[день] = (*рыбы)[день+1]
		} else {
			žuvys[день] = 0
		}
	}
	if born := (*рыбы)[0]; born > 0 {
		žuvys[6] += born
		žuvys[8] += born
	}
	*рыбы = žuvys
}

func day6() uint64 {
	var (
		scanner = bufio.NewScanner(os.Stdin)
		рыбы    [9]uint64
	)
	scanner.Scan()
	for _, t := range strings.Split(scanner.Text(), ",") {
		n, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal("invalid input: ", err)
		}
		рыбы[n]++
	}

	for day := 0; day < 256; day++ {
		ебитес(&рыбы)
	}

	var kiekŽuvų uint64
	for i := range рыбы {
		kiekŽuvų += рыбы[i]
	}
	return kiekŽuvų
}
