package main

import (
	"github.com/jbury/AdventOfCode-2023/util"

	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, err := util.ParseLinesFromFile("input")

	if err != nil {
		panic(err)
	}

	games := GetAllGames(lines)

	maxGame := &Game{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	fmt.Printf("Sum of valid games: %d", SumValidGameIds(games, maxGame))
	fmt.Printf("Power of minimum cubesets: %d", PowerOfGames(games))
}

func SumValidGameIds(games map[string]*Game, maxGame *Game) int {
	sum := 0

	for _, g := range games {
		if maxGame.Red >= g.Red &&
			maxGame.Green >= g.Green &&
			maxGame.Blue >= g.Blue {
			idValue, err := strconv.Atoi(g.Id)
			if err != nil {
				panic(err)
			}
			sum += idValue
		}
	}

	return sum
}

func PowerOfGames(games map[string]*Game) int {
	sumOfPowers := 0

	for _, game := range games {
		sumOfPowers += game.Red * game.Green * game.Blue
	}

	return sumOfPowers
}

func GetAllGames(lines []string) map[string]*Game {
	games := make(map[string]*Game, len(lines))
	for _, l := range lines {
		newGame := gameFromLine(l)

		games[newGame.Id] = newGame
	}

	return games
}

func gameFromLine(line string) *Game {
	idAndRounds := strings.Split(line, ": ")
	fullIdString := strings.TrimSpace(idAndRounds[0])
	roundsString := strings.TrimSpace(idAndRounds[1])

	id := strings.TrimSpace(strings.Split(fullIdString, " ")[1])

	rounds := strings.Split(roundsString, ";")

	newGame := &Game{
		Id:    id,
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	for _, round := range rounds {
		shown := strings.Split(strings.TrimSpace(round), ", ")
		for _, c := range shown {
			colorCount := strings.Split(strings.TrimSpace(c), " ")
			countString := strings.TrimSpace(colorCount[0])
			colorString := strings.TrimSpace(colorCount[1])

			count, err := strconv.Atoi(countString)
			if err != nil {
				panic(fmt.Errorf("tried to Atoi %s, which was before whitespace of %s, which was split by ',' from %s", colorString, colorCount, round))
			}
			newGame.UpdateIfLarger(count, colorString)
		}
	}

	return newGame
}

type Game struct {
	Id    string
	Red   int
	Green int
	Blue  int
}

func (g *Game) UpdateIfLarger(count int, color string) {
	switch color {
	case "red":
		if g.Red < count {
			g.Red = count
		}
	case "green":
		if g.Green < count {
			g.Green = count
		}
	case "blue":
		if g.Blue < count {
			g.Blue = count
		}
	default:
		panic(fmt.Errorf("no clue what I'm supposed to do with the color: %s", color))
	}
}
