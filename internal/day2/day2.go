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

				switch cubeParts[1] {
				case "red":
					if cubeNumber > 12 {
						continue GameLoop
					}
				case "green":
					if cubeNumber > 13 {
						continue GameLoop
					}
				case "blue":
					if cubeNumber > 14 {
						continue GameLoop
					}
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

				switch cubeParts[1] {
				case "red":
                    if cubeNumber > red {
                        red = cubeNumber
                    }
				case "green":
                    if cubeNumber > green {
                        green = cubeNumber
                    }
				case "blue":
                    if cubeNumber > blue {
                        blue = cubeNumber
                    }
				}
			}
		}
		sum += red * green * blue

	}
	return sum
}
