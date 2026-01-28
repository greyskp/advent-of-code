package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/greyskp/advent-of-code/utils"
)

func IdentifyInvalidIds(file string) int {
	data := utils.ReadInputFile(file)
	result := 0
	ranges := strings.SplitSeq(data, ",")

	for ids := range ranges {
		elems := strings.Split(ids, "-")
		start, err := strconv.Atoi(elems[0])
		end, err := strconv.Atoi(elems[1])

		if err != nil {
			log.Fatalf("Error parsing string %v", err)
		}

		for i := start; i <= end; i++ {
			value := strconv.Itoa(i)
			if len(value)%2 == 0 {
				if value[0:len(value)/2] == value[len(value)/2:] {
					result += i
				}
			}
		}
	}

	return result

}

func main() {
	path := "small_entries.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fmt.Println(IdentifyInvalidIds(path))
}
