package main

import (
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

func rotateKnob(file string) int {
	count := 0
	val := 0
	position := initialPosition
	data := strings.SplitSeq(readInputFile(file), "\n")
	for rows := range data {
		val, _ = strconv.Atoi(rows[1:])
		if string(rows[0]) == "L" {
			position -= val
		} else {
			position += val
		}

		if position%100 == 0 {
			count++
		}
	}
	return count
}
