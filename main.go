package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput() string {
	data := os.Args[1]
	return data
}

func main() {
	input := readInput()

	if len(input) == 0 {
		fmt.Println(0)
	} else {
		split := strings.Split(input, " ")
		fmt.Println(len(split))
	}
}