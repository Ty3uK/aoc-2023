package day2

import (
	"bufio"
	"strconv"
	"strings"
)

func Day2(input string) int {
	sum := 0
	scanner := bufio.NewScanner(strings.NewReader(input))

GameLoop:
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ": ")

		gameID, err := strconv.Atoi(game[0][5:])
		if err != nil {
			panic(err)
		}

		for _, subset := range strings.Split(game[1], "; ") {
			for _, cube := range strings.Split(subset, ", ") {
				cubeParts := strings.Split(cube, " ")
				cubeNumber, err := strconv.Atoi(cubeParts[0])
				if err != nil {
					panic(err)
				}

				switch {
				case cubeParts[1] == "red" && cubeNumber > 12:
					fallthrough
				case cubeParts[1] == "green" && cubeNumber > 13:
					fallthrough
				case cubeParts[1] == "blue" && cubeNumber > 14:
					continue GameLoop
				}
			}
		}
		sum += gameID

	}
	return sum
}

func Day2_2(input string) int {
	sum := 0
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ": ")

		red := 0
		green := 0
		blue := 0
		for _, subset := range strings.Split(game[1], "; ") {
			for _, cube := range strings.Split(subset, ", ") {
				cubeParts := strings.Split(cube, " ")
				cubeNumber, err := strconv.Atoi(cubeParts[0])
				if err != nil {
					panic(err)
				}

				switch {
				case cubeParts[1] == "red" && cubeNumber > red:
					red = cubeNumber
				case cubeParts[1] == "green" && cubeNumber > green:
					green = cubeNumber
				case cubeParts[1] == "blue" && cubeNumber > blue:
					blue = cubeNumber
				}
			}
		}
		sum += red * green * blue

	}
	return sum
}
