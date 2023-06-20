package main

import (
	loadData "advent-of-code/utils/loadData"
	"fmt"
	"log"
)

func main() {
	// load file input
	input := loadData.LoadData("./input.txt")
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
