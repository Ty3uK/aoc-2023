package day1

import (
	"bufio"
	"strings"
	"unicode"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parseLine(line string) int {
	firstDigit := -1
	lastDigit := -1

	for _, char := range line {
		if unicode.IsDigit(char) {
			lastDigit = int(char - '0')
			if firstDigit == -1 {
				firstDigit = lastDigit
			}
		}
	}

	return firstDigit*10 + lastDigit
}

func parseLine2(line string) int {
	firstDigit := -1
	lastDigit := -1

	firstDigitIndex := -1
	lastDigitIndex := -1

	for i, char := range line {
		if unicode.IsDigit(char) {
			lastDigit = int(char - '0')
			lastDigitIndex = i
			if firstDigit == -1 {
				firstDigit = lastDigit
				firstDigitIndex = i
			}
		}
	}

	for word, digit := range digits {
		fIdx := strings.Index(line, word)
		lIdx := strings.LastIndex(line, word)

		if fIdx > -1 && fIdx < firstDigitIndex {
			firstDigit = digit
			firstDigitIndex = fIdx
		}
		if lIdx > -1 && lIdx > lastDigitIndex {
			lastDigit = digit
			lastDigitIndex = lIdx
		}
	}

	return firstDigit*10 + lastDigit
}

func Day1(input string) int {
	result := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		result += parseLine(scanner.Text())
	}
	return result
}

func Day1_2(input string) int {
	result := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		result += parseLine2(scanner.Text())
	}
	return result
}
