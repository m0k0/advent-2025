package day05

import (
	"errors"
	"fmt"
	"m0k0/advent-2025/common"
)

func Solve(advent *common.AdventSetup) (string, error) {

	inputFile, err := advent.OpenInput()
	if err != nil {
		return "", fmt.Errorf("day %d error: \n%w", advent.Day, err)
	}
	defer inputFile.Close()

	return "", errors.New("no solution available")
}
