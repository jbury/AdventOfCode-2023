package main

import (
	"github.com/jbury/AdventOfCode-2023/util"
)

func main() {
	lines, err := util.ParseLinesFromFile("input")

	if err != nil {
		panic(err)
	}

	SumPartNumbers(lines)
}

func SumPartNumbers(schematicLines []string) int {
	return 0
}
