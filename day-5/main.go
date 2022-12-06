package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./data/input.txt")

	lines := strings.Split(string(data), "\n")

	columns := getNumberOfColumns(lines)
	maxStackSize := getMaxStackSize(lines)

	stacksPartOne := initializeStacks(lines, columns, maxStackSize)
	stacksPartTwo := initializeStacks(lines, columns, maxStackSize)

	for i := maxStackSize + 2; i < len(lines); i++ {
		var numMoves, from, to int
		fmt.Fscanf(strings.NewReader(lines[i]), "move %d from %d to %d", &numMoves, &from, &to)

		stacksPartTwo[to-1].PushN(stacksPartTwo[from-1].PopN(numMoves))

		for j := 0; j < numMoves; j++ {
			stacksPartOne[to-1].Push(stacksPartOne[from-1].Pop())
		}

	}

	fmt.Print("Part One - ")
	for i := 0; i < len(stacksPartOne); i++ {
		s := stacksPartOne[i]
		fmt.Print(string(s.Pop()))
	}
	fmt.Println("")
	fmt.Print("Part Two - ")

	for i := 0; i < len(stacksPartTwo); i++ {
		s := stacksPartTwo[i]
		fmt.Print(string(s.Pop()))
	}
	fmt.Println("")

}

func initializeStacks(lines []string, numStacks, maxStackSize int) map[int]*stack {
	stacks := make(map[int]*stack)
	colIndex := 1
	for i := 0; i < numStacks; i++ {
		s := &stack{
			members: &[]byte{},
		}

		// Controls the row to draw from, if next == -1 or empty char then continue
		row := maxStackSize - 1

		var next byte
		for row > -1 {
			next = lines[row][colIndex]
			if next == ' ' {
				break
			}
			s.Push(next)
			row--
		}

		stacks[i] = s

		colIndex += 4
	}
	return stacks
}

func getNumberOfColumns(lines []string) int {
	for _, line := range lines {
		if line[1] == '1' {
			split := strings.Split(line, " ")

			columns, _ := strconv.Atoi(split[len(split)-2])

			return columns
		}
	}
	return -1
}

func getMaxStackSize(lines []string) int {
	for lineIndex, line := range lines {
		if line[1] == '1' {
			return lineIndex
		}
	}
	return -1
}

type stack struct {
	members *[]byte
}

func (s *stack) Pop() byte {
	val := (*s.members)[len(*s.members)-1]
	*s.members = (*s.members)[:len(*s.members)-1]
	return val
}

func (s *stack) PopN(n int) []byte {
	val := (*s.members)[len(*s.members)-n:]
	*s.members = (*s.members)[:len(*s.members)-n]
	return val
}

func (s *stack) Push(r byte) {
	*s.members = append(*s.members, r)
}

func (s *stack) PushN(r []byte) {
	*s.members = append(*s.members, r...)
}

func (s *stack) Print() {
	var memberStr []string
	for _, mem := range *s.members {
		memberStr = append(memberStr, string(mem))
	}

	fmt.Printf("Stack: %s\n", strings.Join(memberStr, ", "))
}
