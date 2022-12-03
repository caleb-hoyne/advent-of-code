package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./data/input.txt")
	if err != nil {
		log.Fatal("unable to read file: %w", err)
	}

	bundles := strings.Split(string(data), "\n")

    partOneTotalPriority := partOne(bundles)
    partTwoTotalPriority := partTwo(bundles)

	fmt.Printf("Part One - The total priority is: %d\n", partOneTotalPriority)
    fmt.Printf("Part Two - The total priority is: %d\n", partTwoTotalPriority)
}

func partTwo(bundles []string) int32 {
    var totalPriority int32
    for i := 0; i < len(bundles) - 3; i+=3 {
        bundleOne := compartmentAsMap(bundles[i])
        bundleTwo := compartmentAsMap(bundles[i + 1])
        bundleThree := compartmentAsMap(bundles[i + 2])

        for item := range bundleOne {
            _, ok := bundleTwo[item]
            if !ok {
                continue
            }
            _, ok = bundleThree[item]
            if !ok {
                continue
            }

            totalPriority += getPriority(item)
        }
    }
    return totalPriority
}

func partOne(bundles []string) int32 {
    var totalPriority int32
    for _, bundle := range bundles {

        compartmentOne, compartmentTwo := bundle[:len(bundle)/2], bundle[len(bundle)/2:]
        mapOne, mapTwo := compartmentAsMap(compartmentOne), compartmentAsMap(compartmentTwo)

        for item := range mapOne {
            if _, ok := mapTwo[item]; ok {
                totalPriority += getPriority(item)
            }
        }
    }
    return totalPriority
}

func isLowerCase(char int32) bool {
	return char >= 97 && char <= 122
}

func compartmentAsMap(s string) map[int32]bool {
	compartmentMap := make(map[int32]bool)
	for _, char := range s {
		compartmentMap[char] = true
	}
	return compartmentMap
}

func getPriority(item int32) int32 {
    if isLowerCase(item) {
        return item - 96
    }
    return item - 38
}
