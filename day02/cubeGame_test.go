package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jbury/AdventOfCode-2023/day02"
)

var _ = Describe("CubeGame", func() {
	It("Passes the first example test case", func() {
		testInput := []string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		}

		games := main.GetAllGames(testInput)

		testMaxGame := &main.Game{
			Red:   12,
			Green: 13,
			Blue:  14,
		}

		Expect(8).To(Equal(main.SumValidGameIds(games, testMaxGame)))
	})

	It("Passes the second example test case", func() {
		testInput := []string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		}

		games := main.GetAllGames(testInput)

		Expect(2286).To(Equal(main.PowerOfGames(games)))
	})

})
