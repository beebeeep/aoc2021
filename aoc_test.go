package main

import (
	"fmt"
	"os"
	"testing"
)

func Test__findByCriteria(t *testing.T) {
	BIT_SIZE = 5
	values := []int{
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}
	if findByCriteria(values, true, 0) != 23 {
		t.Error("wrong oxygen")
	}
	if findByCriteria(values, false, 0) != 10 {
		t.Error("wrong co2")
	}

}

func Test_day4(t *testing.T) {
	var err error
	os.Stdin, err = os.Open("day4.in")
	if err != nil {
		t.Error(err)
	}
	day4()
}

func Test_day5(t *testing.T) {
	var err error
	os.Stdin, err = os.Open("day5.test-in")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(day5())
}

func Test_day6(t *testing.T) {
	f := []uint8{3, 4, 3, 1, 2}
	for i := 0; i < 18; i++ {
		ебитес(&f)
		fmt.Println(i, f)
	}
}

func Test_win(t *testing.T) {
	b := &board{
		{{number: 0}, {number: 0}, {number: 0}, {number: 0}, {number: 0}},
		{{number: 0}, {number: 0}, {number: 0}, {number: 0}, {number: 0}},
		{{number: 1}, {number: 1}, {number: 1}, {number: 1}, {number: 1}},
		{{number: 0}, {number: 0}, {number: 0}, {number: 0}, {number: 0}},
		{{number: 0}, {number: 0}, {number: 0}, {number: 0}, {number: 0}},
	}
	mark([]*board{b}, 1)
	if !win(b) {
		t.Fail()
	}
	b = &board{
		{{number: 0}, {number: 0}, {number: 1}, {number: 0}, {number: 0}},
		{{number: 0}, {number: 0}, {number: 1}, {number: 0}, {number: 0}},
		{{number: 0}, {number: 0}, {number: 1}, {number: 0}, {number: 0}},
		{{number: 0}, {number: 0}, {number: 1}, {number: 0}, {number: 0}},
		{{number: 0}, {number: 0}, {number: 1}, {number: 0}, {number: 0}},
	}
	mark([]*board{b}, 1)
	if !win(b) {
		t.Fail()
	}
	b = &board{
		{{number: 0}, {number: 0}, {number: 0}, {number: 0}, {number: 0}},
		{{number: 0}, {number: 0}, {number: 0}, {number: 0}, {number: 0}},
		{{number: 0}, {number: 0}, {number: 0}, {number: 0}, {number: 0}},
		{{number: 0}, {number: 0}, {number: 0}, {number: 0}, {number: 0}},
		{{number: 0}, {number: 0}, {number: 0}, {number: 0}, {number: 0}},
	}
	mark([]*board{b}, 1)
	if win(b) {
		t.Fail()
	}
}
