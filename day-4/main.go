package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

type overlapFunc func(lowerOne, lowerTwo, upperOne, upperTwo int) bool

func main() {
    data, _ := os.ReadFile("./data/input.txt")

    pairs := strings.Split(string(data), "\n")

    fmt.Printf("Part One - The total overlaps is: %d\n", overlapProcessor(pairs, totalOverlap))
    fmt.Printf("Part Two - The total overlaps is: %d\n", overlapProcessor(pairs, partialOverlap))

}
func overlapProcessor(pairs []string, isOverlapping overlapFunc) int {
    var totalOverlaps int
    for _, pair := range pairs {
        pairData := strings.Split(pair, ",")

        lowerOne, upperOne := getAssignment(pairData[0])
        lowerTwo, upperTwo := getAssignment(pairData[1])

        if isOverlapping(lowerOne, lowerTwo, upperOne, upperTwo) {
            totalOverlaps ++
        }

    }
    return totalOverlaps
}

func getAssignment(data string) (int, int) {
    res := strings.Split(data, "-")

    lower, _ := strconv.Atoi(res[0])
    upper, _ := strconv.Atoi(res[1])

    return lower, upper
}

func totalOverlap(lowerOne, lowerTwo, upperOne, upperTwo int) bool {
    if lowerOne >= lowerTwo && upperOne <= upperTwo {
        return true
    }
    if lowerTwo >= lowerOne && upperTwo <= upperOne {
        return true
    }
    return false
}

func partialOverlap(lowerOne, lowerTwo, upperOne, upperTwo int) bool {
    if lowerOne <= upperTwo && upperOne >= lowerTwo {
        return true
    }
    
    return false
}
