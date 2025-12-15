package day04

import (
	"fmt"
	"m0k0/advent-2025/common"
)

func Solve(advent *common.AdventSetup) (string, error) {

	inputFile, err := advent.OpenInput()
	if err != nil {
		return "", fmt.Errorf("day 2 error: \n%w", err)
	}
	defer inputFile.Close()

	return "", nil
}
