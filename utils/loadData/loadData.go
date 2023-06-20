package utils

import (
	"advent-of-code/utils/check"
	"os"
	"strings"
)

func LoadData(path string) []string {
	file, err := os.ReadFile(path)
	utils.Check(err)
	return strings.Split(string(file), "")
}
