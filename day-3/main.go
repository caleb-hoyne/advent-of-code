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

	var totalPriority int32
	for _, bundle := range bundles {

		compartmentOne, compartmentTwo := bundle[:len(bundle)/2], bundle[len(bundle)/2:]
		mapOne, mapTwo := compartmentAsMap(compartmentOne), compartmentAsMap(compartmentTwo)

		for item := range mapOne {
			if _, ok := mapTwo[item]; ok {
				if isLowerCase(item) {
					totalPriority += item - 96
					continue
				}
				totalPriority += item - 38
			}
		}
	}

	fmt.Printf("Part One - The total priority is: %d\n", totalPriority)

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
