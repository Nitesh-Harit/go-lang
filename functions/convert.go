package main

import (
	"fmt"
)

func main() {
	bestFinish := championshipFinishes(12, 5, 6, 3, 3, 3, 1)

	fmt.Println(bestFinish)
}

func championshipFinishes(Finishes ...int) int {
	best := Finishes[0]

	for _, i := range Finishes {
		if i < best {
			best = i
		}
	}
	return best
}
