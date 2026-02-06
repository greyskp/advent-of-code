package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/greyskp/advent-of-code/utils"
)

func PrepareData(filename string) ([][]int, []string) {
	data := utils.ReadInputFile(filename)
	rows := strings.Split(data, "\n")
	lines := make([][]int, 0)
	operations := make([]string, 0)

	for i := 0; i < len(rows); i++ {
		cleanExtraSpaces := strings.TrimSpace(rows[i])
		line := strings.Split(cleanExtraSpaces, " ")

		if i == len(rows)-1 {
			for _, symbol := range line {
				if symbol != "" {
					operations = append(operations, symbol)
				}
			}

		} else {
			lines = append(lines, parseIntList(line))
		}
	}

	return lines, operations
}

func parseIntList(line []string) []int {
	convertedLine := make([]int, 0)
	for _, elem := range line {
		if elem != "" {
			converted, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatalf("Error when parsing an int %v", err)
			}
			convertedLine = append(convertedLine, converted)
		}

	}

	return convertedLine
}

func MathSolver(filename string) int {
	result := 0
	lines, operations := PrepareData(filename)
	if len(lines[0]) != len(operations) {
		log.Fatalf("Size error, size of a line %d, size of the operations %d", len(lines[0]), len(operations))
	}
	numberOfCalculations := len(operations)
	numberOfElements := len(lines)

	for i := 0; i < numberOfCalculations; i++ {
		currentCount := lines[0][i]

		for j := 1; j < numberOfElements; j++ {
			if operations[i] == "+" {
				currentCount += lines[j][i]
			} else {
				currentCount *= lines[j][i]
			}
		}

		result += currentCount

	}

	return result
}

func main() {
	path := "entries.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fmt.Println(MathSolver(path))
}
