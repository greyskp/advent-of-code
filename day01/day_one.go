package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const initialPosition = 50

func readInputFile(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func rotateKnob(file string) (int, int) {
	exactZeros := 0
	passingZeros := 0
	currentSign := 1
	val := 0
	position := initialPosition
	prev := 0
	data := strings.SplitSeq(readInputFile(file), "\n")
	for rows := range data {
		val, _ = strconv.Atoi(rows[1:])

		passingZeros += val / 100
		fmt.Println("multiple laps: ", val/100)
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
			fmt.Println("test passing: ", position, prev)
		} else {
			if position%100 == 0 {
				exactZeros++
			}
		}
		position = position % 100
	}
	fmt.Println("TEST FIN: ", exactZeros, passingZeros)
	passingZeros += exactZeros
	return exactZeros, passingZeros
}
