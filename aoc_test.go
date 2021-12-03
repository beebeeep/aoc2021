package main

import (
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
