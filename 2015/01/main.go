package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// error checker
func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func main() {
	// load file input
	file, err := os.ReadFile("./input.txt")
	check(err)
	input := strings.Split(string(file), "")
	floor := 0

	for k, char := range input {
		if char == "(" {
			floor++
		}
		if char == ")" {
			floor--

			if floor == -1 {
				log.Fatal("pos: ", k+1)
			}
		}
	}

	fmt.Println("floor: ", floor)
}
