package main_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jbury/AdventOfCode-2023/day03"
)

var _ = Describe("EngineSchematic", func() {
	It("Passes the first test case", func() {
		// Oh god how do I make this not so trash looking??
		schematic := strings.Split(
			`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, "\n")

		Expect(4361).To(Equal(main.SumPartNumbers(schematic)))
	})
})
