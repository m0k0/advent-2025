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
	rangeIndex := 0
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

		fmt.Printf("min: %d; max: %d\n", rangeMin, rangeMax)
	}

	return "", nil
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
