package day1

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func SumCalibrationValue() int {
	scanner := bufio.NewScanner(os.Stdin)

	numberArray := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var total int = 0

	for scanner.Scan() {
		line := scanner.Text()

		var first string = ""
		var firstPos int = math.MaxInt
		var last string = ""
		var lastPos int = math.MinInt

		for pos, char := range line {
			if unicode.IsDigit(char) {
				if first == "" {
					first = string(char)
					last = string(char)

					firstPos = pos
					lastPos = pos
				} else {
					last = string(char)
					lastPos = pos
				}
			}
		}

		for pos, num := range numberArray {
			firstIdx := strings.Index(line, num)

			if firstIdx != -1 && firstIdx < firstPos {
				first = strconv.Itoa(pos + 1)
				firstPos = firstIdx
			}

			lastIdx := strings.LastIndex(line, num)

			if lastIdx != -1 && lastIdx > lastPos {
				last = strconv.Itoa(pos + 1)
				lastPos = lastIdx
			}
		}

		var numStr string = ""
		numStr += first
		numStr += last

		number, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		total += number
	}
	return total
}
