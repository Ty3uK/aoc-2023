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
	runes := []rune(line)
	firstDigit := -1
	lastDigit := -1
	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		if firstDigit == -1 {
			if unicode.IsDigit(runes[i]) {
				firstDigit = int(runes[i] - '0')
			}
		}
		if lastDigit == -1 {
			if unicode.IsDigit(runes[j]) {
				lastDigit = int(runes[j] - '0')
			}
		}
		if firstDigit != -1 && lastDigit != -1 {
			break
		}
	}
	return firstDigit*10 + lastDigit
}

func parseLine2(line string) int {
	runes := []rune(line)
	firstDigit := -1
	lastDigit := -1
	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		if firstDigit == -1 {
			if unicode.IsDigit(runes[i]) {
				firstDigit = int(runes[i] - '0')
			} else {
				for word, digit := range digits {
					if i+len(word) >= len(runes) || string(runes[i:i+len(word)]) != word {
						continue
					}
					firstDigit = digit
				}
			}
		}
		if lastDigit == -1 {
			if unicode.IsDigit(runes[j]) {
				lastDigit = int(runes[j] - '0')
			} else {
				for word, digit := range digits {
					if j-len(word) < 0 || string(runes[j-len(word)+1:j+1]) != word {
						continue
					}
					lastDigit = digit
				}
			}
		}
		if firstDigit != -1 && lastDigit != -1 {
			break
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
