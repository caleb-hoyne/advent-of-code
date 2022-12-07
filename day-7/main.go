package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./data/input.txt")

	input := strings.Split(string(data), "\n")

	parent := createFS(input)

	fmt.Printf("Part one - %d\n", partOne(parent))
	fmt.Printf("Part two - %d\n", partTwo(parent))
}

func partTwo(parent *Directory) int {
	totalCurrentStored := parent.GetSize()
	var directorySizes []int
	subDirectories := []*Directory{parent}
	for len(subDirectories) > 0 {
		subDir := subDirectories[0]
		directorySizes = append(directorySizes, subDir.GetSize())

		for key, subSubDir := range subDir.subDirectories {
			if key == ".." {
				continue
			}
			subDirectories = append(subDirectories, subSubDir)
		}
		subDirectories = subDirectories[1:]
	}

	sort.Ints(directorySizes)

	totalSpace := 70000000
	neededSpace := 30000000
	maxAllowedSpace := totalSpace - neededSpace

	for _, dirSize := range directorySizes {
		if (totalCurrentStored - dirSize) <= maxAllowedSpace {
			return dirSize
		}
	}
	return -1
}

func partOne(parent *Directory) int {
	var totalDirSize int
	subDirectories := []*Directory{parent}
	for len(subDirectories) > 0 {
		subDir := subDirectories[0]
		if subDir.GetSize() <= 100000 {
			totalDirSize += subDir.GetSize()
		}

		for key, subSubDir := range subDir.subDirectories {
			if key == ".." {
				continue
			}
			subDirectories = append(subDirectories, subSubDir)
		}
		subDirectories = subDirectories[1:]
	}
	return totalDirSize
}

func createFS(input []string) *Directory {
	cur := &Directory{
		subDirectories: map[string]*Directory{},
	}
	parent := cur

	lineInd := 0
	for lineInd < len(input)-1 {
		lineInd++
		line := input[lineInd]
		switch strings.Split(line, " ")[0] {
		case "$":
			if getCommand(line) == "cd" {
				dir := getCmdDirectory(line)
				newDir, ok := cur.subDirectories[dir]
				if ok {
					cur = newDir
					continue
				}
				parentTmp := cur
				cur = &Directory{
					subDirectories: map[string]*Directory{
						"..": parentTmp,
					},
				}
			}
		case "dir":
			cur.subDirectories[getLSDirectory(line)] = &Directory{
				subDirectories: map[string]*Directory{
					"..": cur,
				},
			}
		default:
			cur.files = append(cur.files, file{
				name: getFileName(line),
				size: getFilSize(line),
			})
		}
	}
	return parent
}

func getFileName(line string) string {
	return strings.Split(line, " ")[1]
}

func getFilSize(line string) int {
	spl := strings.Split(line, " ")[0]
	size, err := strconv.Atoi(spl)
	if err != nil {
		panic(err)
	}
	return size
}

func getLSDirectory(line string) string {
	return strings.Split(line, " ")[1]
}

func getCmdDirectory(line string) string {
	return strings.Split(line, " ")[2]
}

func getCommand(line string) string {
	return strings.Split(line, " ")[1]
}

type Directory struct {
	files          []file
	subDirectories map[string]*Directory
}

func (d *Directory) GetSize() int {
	var dirSize int
	for _, f := range d.files {
		dirSize += f.size
	}
	for k, subDir := range d.subDirectories {
		if k == ".." {
			continue
		}
		dirSize += subDir.GetSize()
	}
	return dirSize
}

type file struct {
	name string
	size int
}
