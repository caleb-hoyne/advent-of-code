package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
    "sort"
)

func main() {
	data, err := os.ReadFile("./data/input.txt")
	if err != nil {
		log.Fatal("unable to read file: %w", err)
	}

	calorieLists := strings.Split(string(data), "\n\n")


    numReindeer := len(calorieLists)
    calIntList := make([]int, 0, numReindeer)
	for _, calList := range calorieLists {
        calIntList = append(calIntList, getCalories(calList))
	}

    sort.Ints(calIntList)

    totalMax := calIntList[numReindeer - 1] + calIntList[numReindeer - 2] + calIntList[numReindeer - 3]

	fmt.Printf("The total max calories is %d\n", totalMax)
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
