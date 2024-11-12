package main

import (
	"fmt"
	"iter"
	"strings"

	"github.com/RaphaelPour/stellar/input"
)

func slidingWindowSeq[T any](slice []T, size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for i := 0; i < len(slice) && len(slice[i:size]) >= size; i += 1 {
			fmt.Println(slice[i:size])
			if !yield(slice[i:size]) {
				return
			}
		}
	}
}

func part1(file string) {
	in := input.LoadString(file)

	if len(in) != 3 {
		fmt.Printf("expected 3 lines, got %d instead\n", len(in))
		return
	}

	words := strings.Split(strings.TrimPrefix(in[0], "WORDS:"), ",")
	text := strings.Split(in[2], " ")

	sum := 0
	for _, textWord := range text {
		for _, word := range words {
			sum += strings.Count(textWord, word)
		}
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println(" ===[ PART   I ]===")

	for _, file := range []string{
		"input_example",
		"input_example2",
		"input_example3",
		"input_example4",
		"input",
	} {
		part1(file)
	}
}
