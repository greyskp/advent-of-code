package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/greyskp/advent-of-code/utils"
)

var neighbors = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1},           {0, 1},
	{1, -1},  {1, 0},  {1, 1},
}

func ForkliftMap(filename string) int {
	result := 0
	paperRollsMap := utils.ReadInputFile(filename)
	rows := strings.Split(paperRollsMap, "\n")

	for i, row := range rows {
		for j := range row {
			
			count := 0
			
			if rows[i][j] == '@' {

				for _, d := range neighbors {
					ni := i + d[0]
					//fmt.Println(ni)
					if ni < 0 || ni >= len(rows) {
						continue
					}

					nj := j + d[1]
					if nj < 0 || nj >= len(rows[ni]) {
						continue
					}

					if rows[ni][nj] == '@' {
						count++
						if count >= 4 {
							break
						}
					}
				}

				if count < 4 {
					result++ 
				}
			}
			
			

		}
	}

	return result
}

func main() {
	path := "entries.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fmt.Println(ForkliftMap(path))
}
