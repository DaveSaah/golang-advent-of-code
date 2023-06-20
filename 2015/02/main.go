package main

import (
	loadData "advent-of-code/utils/loadData"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parea(l, w, h float64) (int, int) {
	sarea := 2 * ((l * w) + (w * h) + (h * l))
	max := math.Max(l, math.Max(w, h))
	extra := (l * w * h) / max
	lribbon := (2 * (l + w + h)) - (2 * max) + (l * w * h)
	return int(sarea + extra), int(lribbon)
}

func main() {
	input := loadData.LoadData(loadData.DataParams{Path: "./input.txt", Sep: "\n"})
	totalArea := 0
	totalRibbon := 0

	// test := [2]string{"1x1x10", ""}

	for _, dims := range input {
		if dims == "" {
			continue
		}

		dim := strings.Split(dims, "x")
		l, _ := strconv.ParseFloat(dim[0], 64)
		w, _ := strconv.ParseFloat(dim[1], 64)
		h, _ := strconv.ParseFloat(dim[2], 64)
		larea, lribbon := parea(l, w, h)
		totalArea += larea
		totalRibbon += lribbon
	}

	fmt.Println("total area:", totalArea)
	fmt.Println("total ribbon length:", totalRibbon)
}
