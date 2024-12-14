package main

import (
	"testing"
)

func TestEvolve(t *testing.T) {
	var tests = []struct {
		name       string
		testpebble []int
		blinks     int
		want       int // number of pebbles
	}{
		{
			"test1",
			[]int{125, 17},
			3,
			5,
		},
		{
			"test2",
			[]int{0, 1, 10, 99, 999},
			1,
			7,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Evolve(test.testpebble, test.blinks)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
