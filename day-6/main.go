package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./data/input.txt")

	input := strings.Split(string(data), "\n")[0]

	fmt.Printf("Part One - %d\n", solve(4, input))
	fmt.Printf("Part Two - %d\n", solve(14, input))
}

func solve(seqLength int, input string) int {

	currentSeq := make(map[byte]int)
	for i := 0; i < seqLength; i++ {
		currentSeq[input[i]] = currentSeq[input[i]] + 1
	}

	startingInd := seqLength
	for {
		currentChar := input[startingInd]
		prevChar := input[startingInd-seqLength]

		currentSeq[currentChar] = currentSeq[currentChar] + 1
		val := currentSeq[prevChar]
		if val == 1 {
			delete(currentSeq, prevChar)
		} else {
			currentSeq[prevChar] = currentSeq[prevChar] - 1
		}

		if len(currentSeq) == seqLength {
			break
		}

		startingInd++
	}

	return startingInd + 1
}
