package day02

import (
	"bufio"
	"fmt"
	"m0k0/advent-2025/common"
	"strconv"
	"strings"
)

func Solve(advent *common.AdventSetup) (string, error) {

	inputFile, err := advent.OpenInput()
	if err != nil {
		return "", fmt.Errorf("day 2 error: \n%w", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(splitRange)

	const RANGE_SEPARATOR string = "-"
	var rangeIndex int64 = 0
	var invalidIds int64 = 0
	for scanner.Scan() {
		rangeIndex++
		rangeText := scanner.Text()
		rangeText = strings.Trim(rangeText, "\r\n")

		rangeSegments := strings.Split(rangeText, RANGE_SEPARATOR)

		if len(rangeSegments) != 2 {
			return "", fmt.Errorf("invalid range at index %d, too many arguments: '%s'",
				rangeIndex, rangeText)
		}

		rangeMin, err := strconv.ParseInt(rangeSegments[0], 10, 64)
		if err != nil {
			return "", fmt.Errorf("invalid range at index %d, error parsing start of range: '%s'",
				rangeIndex, rangeText)
		}
		rangeMax, err := strconv.ParseInt(rangeSegments[1], 10, 64)
		if err != nil {
			return "", fmt.Errorf("invalid range at index %d, error parsing end of range: '%s'",
				rangeIndex, rangeText)
		}

		invalidIds += sumInvalidIds(rangeMin, rangeMax, advent.VerboseOutput)
	}

	solution := fmt.Sprintf("sum of invalid ids: %d", invalidIds)
	return solution, nil
}

func sumInvalidIds(rangeMin int64, rangeMax int64, verboseOutput bool) int64 {

	var invalidIdSum int64 = 0
	if verboseOutput {
		fmt.Printf("min: %d; max: %d\n", rangeMin, rangeMax)
	}
	for id := rangeMin; id <= rangeMax; id++ {

		if isValidId(id) {
			continue
		}
		if verboseOutput {
			fmt.Printf("\tinvalid id: %d\n", id)
		}
		invalidIdSum += id
	}

	return invalidIdSum
}

func isValidId(id int64) bool {

	idString := fmt.Sprint(id)

	// uneven amount of digits, can't be invalid
	if len(idString)%2 == 1 {
		return true
	}

	firstPart := idString[:len(idString)/2]
	secondPart := idString[len(idString)/2:]

	// not repeating twice, can't be invalid
	if firstPart != secondPart {
		return true
	}

	trimmedIdPart := strings.TrimLeft(firstPart, "0")
	// has leading zeroes, can't be invalid
	if len(trimmedIdPart) != len(firstPart) {
		return true
	}

	// assume invalid
	return false
}

func splitRange(data []byte, atEOF bool) (advance int, token []byte, err error) {

	const RANGE_DELIMITER string = ","

	// end of input, no data
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// delim found, return data before delim
	delimiterIndex := strings.Index(string(data), RANGE_DELIMITER)
	if delimiterIndex > -1 {
		return delimiterIndex + 1, data[0:delimiterIndex], nil
	}

	// end of input, no delim, return everything
	if atEOF {
		return len(data), data, nil
	}

	return
}
