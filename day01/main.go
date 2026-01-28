package main

import (
	"fmt"
	"os"
)

func main() {
	path := "small_entries.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	exact, passing := rotateKnob(path)
	fmt.Println(exact, passing)
}
