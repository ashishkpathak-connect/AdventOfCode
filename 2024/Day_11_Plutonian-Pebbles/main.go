package main

import (
	"fmt"
	"strconv"
)

const (
	zerovalue = 1
	oddvalue  = 2024
)

func Split(elem int) []int {
	elemstr := fmt.Sprintf("%v", elem)
	elemblocksize := len(elemstr) / 2
	leftstr := elemstr[0:elemblocksize]
	rightstr := elemstr[elemblocksize:]
	left, _ := strconv.Atoi(leftstr)
	right, _ := strconv.Atoi(rightstr)
	return []int{left, right}
}

func Evolve(pebbles []int, blinks int) int {
	pebblesmap := make(map[int]int)

	for _, pebble := range pebbles {
		pebblesmap[pebble]++
	}

	for i := 0; i < blinks; i++ {
		newpebblesmap := make(map[int]int)

		for pebble, count := range pebblesmap {
			pebblestr := fmt.Sprintf("%v", pebble)
			if len(pebblestr)%2 != 0 { // odd length
				switch pebble {
				case 0:
					newpebblesmap[zerovalue] += count
				default:
					newpebblesmap[pebble*oddvalue] += count
				}

			} else {
				splitpebble := Split(pebble)
				newpebblesmap[splitpebble[0]] += count
				newpebblesmap[splitpebble[1]] += count
			}
		}
		pebblesmap = newpebblesmap
	}
	var totalpebbles int
	for _, count := range pebblesmap {
		totalpebbles += count
	}
	return totalpebbles
}

func main() {
	pebbles := []int{0, 37551, 469, 63, 1, 791606, 2065, 9983586}
	fmt.Println(Evolve(pebbles, 75))

}
