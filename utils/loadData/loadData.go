package utils

import (
	check "advent-of-code/utils/check"
	"os"
	"strings"
)

type DataParams struct {
	Path string
	Sep  string
}

func LoadData(params DataParams) []string {
	file, err := os.ReadFile(params.Path)
	check.Check(err)
	return strings.Split(string(file), params.Sep)
}
