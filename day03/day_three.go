package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/greyskp/advent-of-code/utils"
)

func HighestJoltage(file string) int {
	res := 0
	data := utils.ReadInputFile(file)
	rows := strings.Split(data, "\n")

	for _, values := range rows {
		first := 0
		second := 0
		valueInt := 0

		for i, elem := range values {
			valueInt, _ = strconv.Atoi(string(elem))
			if first == 0 || (valueInt > first && i < len(values)-1) {
				first = valueInt
				second = 0
			} else {
				if valueInt > second {
					second = valueInt
				}
			}
		}
		resRow, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(second))
		res += resRow
	}

	return res
}

func main() {
	path := "entries.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fmt.Println(HighestJoltage(path))
}
