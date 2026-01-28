package main

import (
	"strconv"
	"strings"

	"github.com/greyskp/advent-of-code/utils"
)

const initialPosition = 50

func rotateKnob(file string) (int, int) {
	exactZeros := 0
	passingZeros := 0
	currentSign := 1
	val := 0
	position := initialPosition
	prev := 0
	data := strings.SplitSeq(utils.ReadInputFile(file), "\n")
	for rows := range data {
		val, _ = strconv.Atoi(rows[1:])

		passingZeros += val / 100
		val = val % 100

		if string(rows[0]) == "L" {
			currentSign = -1
		} else {
			currentSign = 1
		}

		prev = position
		position = position + currentSign*val
		if position%100 != 0 && ((position*prev < 0) || (position > 100 || position < -100)) {
			passingZeros++
		} else {
			if position%100 == 0 {
				exactZeros++
			}
		}
		position = position % 100
	}
	passingZeros += exactZeros
	return exactZeros, passingZeros
}
