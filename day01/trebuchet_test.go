package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jbury/AdventOfCode-2023/day01"
)

var _ = Describe("Trebuchet", func() {
	It("Passes the first example test case", func() {
		testValue := []string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet",
		}

		Expect(142).To(Equal(main.SumCalibrationValues(testValue, false)))
	})
	It("Passes the second example test case", func() {
		testValue := []string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		}

		Expect(281).To(Equal(main.SumCalibrationValues(testValue, true)))
	})
})
