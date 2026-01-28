package utils

import (
	"log"
	"os"
)

func ReadInputFile(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
