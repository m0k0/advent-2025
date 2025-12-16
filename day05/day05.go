package day05

import (
	"bufio"
	"errors"
	"fmt"
	"m0k0/advent-2025/common"
	"strconv"
	"strings"
)

type NumberRange struct {
	start int64
	end   int64
}

func (numberRange *NumberRange) IsInRange(number int64) bool {
	inRange := number >= numberRange.start && number <= numberRange.end
	return inRange
}

func createRange(spec string) (*NumberRange, error) {

	specParts := strings.Split(spec, "-")
	if len(specParts) != 2 {
		return nil, errors.New("range spec does not have exactly 2 parts")
	}

	start, err := strconv.ParseInt(specParts[0], 10, 64)
	if err != nil {
		return nil, errors.New("failed to parse start of range")
	}

	end, err := strconv.ParseInt(specParts[1], 10, 64)
	if err != nil {
		return nil, errors.New("failed to parse end of range")
	}

	numberRange := NumberRange{
		start: start,
		end:   end,
	}

	return &numberRange, nil
}

func Solve(advent *common.AdventSetup) (string, error) {

	logger := common.Logger{
		Verbose: advent.VerboseOutput,
	}

	inputFile, err := advent.OpenInput()
	if err != nil {
		return "", fmt.Errorf("day %d error: \n%w", advent.Day, err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	// read ranges
	freshRanges := common.LinkedList[*NumberRange]{}

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++

		line := scanner.Text()

		if line == "" {
			break // end of range specification
		}

		numberRange, err := createRange(line)
		if err != nil {
			return "", fmt.Errorf("invalid range '%s' on line %d: %w", line, lineNumber, err)
		}

		freshRanges.Add(numberRange)

	}

	// evaluate ingredients

	freshCount := 0
	for scanner.Scan() {
		lineNumber++

		line := scanner.Text()

		id, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return "", fmt.Errorf("invalid id '%s' on line %d", line, lineNumber)
		}

		logger.PrintVerboseF("\ningredient %d\n", id)
		isFresh := false
		for numberRange := range freshRanges.Iterate() {
			if numberRange.IsInRange(id) {
				logger.PrintVerboseFD("is in range %d-%d\n", 1, numberRange.start, numberRange.end)
				isFresh = true
				break
			}
		}

		if isFresh {
			freshCount++
		}
	}

	solution := fmt.Sprintf("total number of fresh ingredients: %d", freshCount)
	return solution, nil
}
