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

func (numberRange *NumberRange) HasInRange(number int64) bool {
	inRange := number >= numberRange.start && number <= numberRange.end
	return inRange
}
func (a *NumberRange) Merge(b *NumberRange) (*NumberRange, error) {

	if a.HasInRange(b.start) &&
		a.HasInRange(b.end) {
		// a fully contains b
		return a, nil
	} else if b.HasInRange(a.start) &&
		b.HasInRange(a.end) {
		// b fully cotains a
		return b, nil
	} else if a.HasInRange(b.start) ||
		b.HasInRange(a.end) {
		new := &NumberRange{
			start: a.start,
			end:   b.end,
		}
		return new, nil
	} else if a.HasInRange(b.end) ||
		b.HasInRange(a.start) {
		new := &NumberRange{
			start: b.start,
			end:   a.end,
		}
		return new, nil
	}
	return nil, errors.New("ranges can't be merged")
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

	var freshCount int64
	if advent.Variant == "part1" {
		// evaluate ingredients
		freshCount, err = countFreshIngredientsFromAvailableIds(scanner, lineNumber, freshRanges, logger)
		if err != nil {
			return "", err
		}

	} else {
		// evaluate ingredients
		optimiseRanges(freshRanges, logger)
		freshCount = countFreshIngredientsFromRanges(freshRanges, logger)
	}

	solution := fmt.Sprintf("total number of fresh ingredients: %d", freshCount)
	return solution, nil
}

func countFreshIngredientsFromRanges(freshRanges common.LinkedList[*NumberRange], logger common.Logger) int64 {

	var totalFreshCount int64 = 0
	for currentRange := range freshRanges.Values() {
		freshCount := currentRange.end - currentRange.start + 1
		logger.PrintVerboseF("range %d-%d: %d fresh ingredients\n", currentRange.start, currentRange.end, freshCount)
		totalFreshCount += freshCount
	}
	return totalFreshCount
}
func optimiseRanges(freshRanges common.LinkedList[*NumberRange], logger common.Logger) {

	//optimisedRanges := common.LinkedList[*NumberRange]{}

	iterationCount := 0

	for {
		iterationCount++
		logger.PrintVerboseF("optimiseRanges iteration %d\n", iterationCount)
		mergeCount := 0
		for entryA := range freshRanges.Entries() {
			currentRange := entryA.Value
			logger.PrintVerboseF("range %d-%d:\n", currentRange.start, currentRange.end)

			for entryB := range freshRanges.Entries() {
				optimisedRange := entryB.Value
				if optimisedRange == nil || currentRange == optimisedRange {
					continue
				}

				newRange, err := optimisedRange.Merge(currentRange)
				if err == nil {
					entryA.Remove()
					entryB.Value = newRange
					mergeCount++
					logger.PrintVerboseFD("merged %d-%d with %d-%d: %d-%d\n", 1,
						currentRange.start, currentRange.end,
						optimisedRange.start, optimisedRange.end,
						newRange.start, newRange.end)
					break
				}
			}
		}
		if mergeCount == 0 {
			break
		}
	}
}

func countFreshIngredientsFromAvailableIds(scanner *bufio.Scanner, lineNumber int, freshRanges common.LinkedList[*NumberRange], logger common.Logger) (int64, error) {
	var freshCount int64 = 0
	for scanner.Scan() {
		lineNumber++

		line := scanner.Text()

		id, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return freshCount, fmt.Errorf("invalid id '%s' on line %d", line, lineNumber)
		}

		logger.PrintVerboseF("\ningredient %d\n", id)
		isFresh := false
		for numberRange := range freshRanges.Values() {
			if numberRange.HasInRange(id) {
				logger.PrintVerboseFD("is in range %d-%d\n", 1, numberRange.start, numberRange.end)
				isFresh = true
				break
			}
		}

		if isFresh {
			freshCount++
		}
	}
	return freshCount, nil
}
