package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	rock     = "A"
	paper    = "B"
	scissors = "C"

	lose = "X"
	draw = "Y"
	win  = "Z"
)

var moveAsPoints = map[string]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

// map of opponent moves to a map of outcomes to the move required to achieve them
var winningMap = map[string]map[string]string{
	rock: {
		win:  paper,
		lose: scissors,
	},
	paper: {
		win:  scissors,
		lose: rock,
	},
	scissors: {
		win:  rock,
		lose: paper,
	},
}

func main() {
	data, err := os.ReadFile("./data/input.txt")
	if err != nil {
		log.Fatal("unable to read file: %w", err)
	}

	games := strings.Split(string(data), "\n")

	var totalScore int
	for _, game := range games {
		moves := strings.Split(game, " ")

		opponentMove, result := moves[0], moves[1]

		var myMove string
		switch result {
		case win:
			totalScore += 6
			myMove = winningMap[opponentMove][win]
		case lose:
			myMove = winningMap[opponentMove][lose]
		case draw:
			totalScore += 3
			myMove = opponentMove
		}
		totalScore += moveAsPoints[myMove]
	}

	fmt.Printf("My total score is: %d\n", totalScore)
}
