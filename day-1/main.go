package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./data/input.txt")
	if err != nil {
		log.Fatal("unable to read file: %w", err)
	}

	calorieLists := strings.Split(string(data), "\n\n")

	var maxCal int
	for _, calList := range calorieLists {
		newCal := getCalories(calList)
		if newCal > maxCal {
			maxCal = newCal
		}
	}

	fmt.Printf("The max calories is %d\n", maxCal)
}

func getCalories(list string) int {

	calories := strings.Split(list, "\n")

	var totalCal int
	for _, calStr := range calories {
		if calStr == "" {
			continue
		}

		cal, err := strconv.Atoi(calStr)
		if err != nil {
			log.Fatalf("%s is not a valid calorie value: %s", calStr, err)
		}
		totalCal += cal
	}

	return totalCal
}
