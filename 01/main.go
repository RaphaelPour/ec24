package main

import (
	"fmt"
	"iter"
	"strings"

	"github.com/RaphaelPour/stellar/input"
)

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func subdivideSeq[T any](in []T, num int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for i := 0; i < len(in); i += num {
			if !yield(in[i:Min(i+num, len(in))]) {
				return
			}
		}
	}
}

func part(inputFile string, num int) {
	sum := 0
	for pair := range subdivideSeq[string](strings.Split(input.LoadString(inputFile)[0], ""), num) {
		team := num
		for _, enemy := range pair {
			switch enemy {
			case "A":
			case "B":
				sum += 1
			case "C":
				sum += 3
			case "D":
				sum += 5
			case "x":
				team -= 1
			}
		}
		if team == 2 {
			sum += 2
		} else if team == 3 {
			sum += 6
		}
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println(" ===[ PART   I ]====")
	part("input", 1)
	fmt.Println(" ===[ PART  II ]====")
	part("input2", 2)
	fmt.Println(" ===[ PART III ]====")
	part("input3", 3)
}
