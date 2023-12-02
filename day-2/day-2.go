package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var colorCountMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func PossibleGameIdSum() int {
	var scanner = bufio.NewScanner(os.Stdin)

	var idSum int = 0

	for scanner.Scan() {
		var line = scanner.Text()
		var gameString = strings.Split(line, ": ")

		var gameId = gameString[0]

		if isGameValid(gameString[1]) {
			var id, err = strconv.Atoi(strings.Split(gameId, "Game ")[1])
			if err != nil {
				log.Fatal(err)
			}
			idSum += id
		}
	}
	return idSum
}

func PowerSetSum() int64 {
	var scanner = bufio.NewScanner(os.Stdin)

	var powerSetSum int64 = 0

	for scanner.Scan() {
		var line = scanner.Text()
		var gameString = strings.Split(line, ": ")

		var gameSet = gameString[1]

		powerSetSum += getGamePowerSet(gameSet)

	}
	return powerSetSum
}

func getGamePowerSet(gameSet string) int64 {
	var gameCubes = strings.Split(gameSet, "; ")

	var colorMap = map[string]int{}

	for _, cubes := range gameCubes {
		var colorCountArray = strings.Split(cubes, ", ")

		for _, colorCount := range colorCountArray {
			var cc = strings.Split(colorCount, " ")

			var count, err = strconv.Atoi(cc[0])
			if err != nil {
				log.Fatal(err)
			}

			var color = cc[1]

			value, ok := colorMap[color]

			if ok {
				colorMap[color] = max(count, value)
			} else {
				colorMap[color] = count
			}
		}
	}

	return int64(colorMap["red"] * colorMap["green"] * colorMap["blue"])
}

func isGameValid(gameSet string) bool {
	var gameCubes = strings.Split(gameSet, "; ")

	for _, cubes := range gameCubes {
		var colorCountArray = strings.Split(cubes, ", ")

		for _, colorCount := range colorCountArray {
			var cc = strings.Split(colorCount, " ")

			var count, err = strconv.Atoi(cc[0])

			if err != nil {
				log.Fatal(err)
			}

			var color = cc[1]

			if count > colorCountMap[color] {
				return false
			}
		}
	}
	return true
}
