package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/greyskp/advent-of-code/utils"
)

func parseRangeAndIngredients(filename string) ([][]int, []int) {
	data := utils.ReadInputFile(filename)
	inventory := strings.Split(data, "\n\n")

	rangeList := parseRanges(strings.Split(inventory[0], "\n"))
	ingredientsList := parseList(strings.Split(inventory[1], "\n"))

	sort.Ints(ingredientsList)

	sort.Slice(rangeList, func(i, j int) bool {
		return rangeList[i][0] < rangeList[j][0]
	})

	return rangeList, ingredientsList
}

func parseList(list []string) []int {
	parsedList := []int{}
	for _, element := range list {
		parsed, err := strconv.Atoi(element)

		if err != nil {
			log.Fatalf("Error when parsing int: %v", err)
		}
		parsedList = append(parsedList, parsed)
	}
	return parsedList
}

func parseRanges(ranges []string) [][]int {
	parsedRanges := [][]int{}

	for _, ingredientRange := range ranges {
		addRange := []int{}
		elem := strings.Split(ingredientRange, "-")

		parsedRangeStart, err := strconv.Atoi(elem[0])
		parsedRangeEnd, err := strconv.Atoi(elem[1])

		if err != nil {
			log.Fatalf("Error when parsing int: %v", err)
		}

		addRange = append(addRange, parsedRangeStart)
		addRange = append(addRange, parsedRangeEnd)

		parsedRanges = append(parsedRanges, addRange)

	}

	return parsedRanges
}

func CountFreshIngredientsFromList(filename string) int {
	freshIngredients := 0

	rangeList, ingredientsList := parseRangeAndIngredients(filename)

	for _, ingredient := range ingredientsList {
		for _, range_ := range rangeList {
			if ingredient >= range_[0] && ingredient <= range_[1] {
				freshIngredients++
				break
			}
		}
	}

	return freshIngredients
}

func CountAllFreshIngredients(filename string) int {
	total := 0
	rangeList, _ := parseRangeAndIngredients(filename)
	latest := 0

	for _, range_ := range rangeList {
		start := max(range_[0], latest+1)

		if start <= range_[1] {
			total += range_[1] - start + 1
			latest = range_[1]
		}
	}

	return total
}

func main() {
	path := "entries.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fmt.Println(CountFreshIngredientsFromList(path))
	fmt.Println(CountAllFreshIngredients(path))
}