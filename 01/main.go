package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	"github.com/RaphaelPour/stellar/iter"
)

func part(inputFile string, num int) {
	sum := 0
	for pair := range iter.PairsSeq(strings.Split(input.LoadString(inputFile)[0], ""), num) {
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
