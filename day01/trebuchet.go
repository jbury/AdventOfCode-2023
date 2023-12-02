package main

import (
	"github.com/jbury/AdventOfCode-2023/util"

	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	lines, err := util.ParseLinesFromFile("input")

	if err != nil {
		panic(err)
	}

	// Set second arg to false to get the first problem's answer
	fmt.Print(SumCalibrationValues(lines, true))
}

// I started writing a more generically powerful way to express this logic while retaining
// the O(n) time complexity of converting a given spelled-digit to it's numerical digit
// equivalent, and it hurt me quite a bit, so I said screw it, I'll just get some practice
// in.
//
// This function replaces all spelled digits substrings (one, two, three, etc.) with their
// numerical counterparts, and strips out all other non-digit characters from a given string
func replaceSpelledDigits(line string) string {
	var newLine strings.Builder
	len := len(line)

	for i := 0; i < len; i++ {
		c := []rune(line)[i]
		switch c {
		// I guess "zero" isn't a valid wordified digit for reasons.
		//		case 'z':
		//			if i+3 < len && line[i+1] == 'e' && line[i+2] == 'r' && line[i+3] == 'o' {
		//				newLine.WriteRune('0')
		//				i += 3
		//				continue
		//			}
		case 'o':
			if i+2 < len && line[i+1] == 'n' && line[i+2] == 'e' {
				newLine.WriteRune('1')
				continue
			}
		case 't':
			if i+2 < len && line[i+1] == 'w' && line[i+2] == 'o' {
				newLine.WriteRune('2')
				continue
			}
			if i+4 < len && line[i+1] == 'h' && line[i+2] == 'r' && line[i+3] == 'e' && line[i+4] == 'e' {
				newLine.WriteRune('3')
				continue
			}
		case 'f':
			if i+3 < len && line[i+1] == 'o' && line[i+2] == 'u' && line[i+3] == 'r' {
				newLine.WriteRune('4')
				continue
			}
			if i+3 < len && line[i+1] == 'i' && line[i+2] == 'v' && line[i+3] == 'e' {
				newLine.WriteRune('5')
				continue
			}
		case 's':
			if i+2 < len && line[i+1] == 'i' && line[i+2] == 'x' {
				newLine.WriteRune('6')
				continue
			}
			if i+4 < len && line[i+1] == 'e' && line[i+2] == 'v' && line[i+3] == 'e' && line[i+4] == 'n' {
				newLine.WriteRune('7')
				continue
			}
		case 'e':
			if i+4 < len && line[i+1] == 'i' && line[i+2] == 'g' && line[i+3] == 'h' && line[i+4] == 't' {
				newLine.WriteRune('8')
				continue
			}
		case 'n':
			if i+3 < len && line[i+1] == 'i' && line[i+2] == 'n' && line[i+3] == 'e' {
				newLine.WriteRune('9')
				continue
			}
		default:
			// Go is shitposting now, right?  This isn't real life, right?
			if unicode.IsDigit(c) {
				newLine.WriteRune(c)
			}
		}
	}

	return newLine.String()
}

func filterInDigits(line string) string {
	var newLine strings.Builder

	for _, c := range line {
		if unicode.IsDigit(c) {
			newLine.WriteRune(c)
		}
	}

	return newLine.String()
}

func SumCalibrationValues(calibrationLines []string, expandWords bool) int {
	sum := 0
	for _, e := range calibrationLines {
		if expandWords {
			sum += extractCalibrationValue(replaceSpelledDigits(e))
		} else {
			sum += extractCalibrationValue(filterInDigits(e))
		}
	}

	return sum
}

func extractCalibrationValue(line string) int {
	val, err := strconv.Atoi(string(line[0]) + string(line[len(line)-1]))

	if err != nil {
		panic(err)
	}

	return val
}
