package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/greyskp/advent-of-code/utils"
)

func HighestJoltagePartTwo(filename string, numberOfSwitches int) int {
	res := 0
	data := utils.ReadInputFile(filename)
	banks := strings.Split(data, "\n")

	// Go through the different battery banks
	for _, bank := range banks {
		switches := make([]int, numberOfSwitches)
		start := 0

		// Go through the different switches in the bank
		for i, joltage := range bank {
			valueJoltage, _ := strconv.Atoi(string(joltage))
			remaining := len(bank) - i

			// Check number of elements left in the row to update starting position
			if numberOfSwitches > remaining {
				start++
			}

			//Find the first switch that can be updated and reset the rest to 0
			for x := start; x < len(switches); x++ {
				if valueJoltage > switches[x] {
					switches[x] = valueJoltage
					clear(switches[x+1:])
					break
				}
			}
		}
		
		var builder strings.Builder
		for _, e := range switches {
			builder.WriteString(strconv.Itoa(e))
		}
		resRow, err := strconv.Atoi(builder.String())

		if err != nil {
			log.Fatalf("Error parsing string %v", err)
		}
		res += resRow
	}

	return res
}

func main() {
	path := "entries.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fmt.Println(HighestJoltagePartTwo(path, 2))
	fmt.Println(HighestJoltagePartTwo(path, 12))
}
