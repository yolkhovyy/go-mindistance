package main

import (
	"fmt"
	"math"
)

func minDistance(a []int) int {
	if len(a) == 0 {
		return -1
	}

	min := math.MaxInt32
	indexMap := make(map[int]int)

	for i, value := range a {
		if prev, found := indexMap[value]; found {
			distance := i - prev
			if distance < min {
				min = distance
			}
		}
		indexMap[value] = i
	}

	if min == math.MaxInt32 {
		return -1
	}

	return min
}

func main() {
	a := []int{3, 2, 0, 8, 2, 3}
	fmt.Printf("%d\n", minDistance(a))
}
