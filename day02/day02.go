package day02

import (
	"bufio"
	"fmt"
	"m0k0/advent-2025/common"
	"math"
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

	validationFunc := isValidId
	if advent.Variant == "part2" {
		validationFunc = isValidId2
	}

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

		invalidIds += sumInvalidIds(rangeMin, rangeMax, validationFunc, advent.VerboseOutput)
	}

	solution := fmt.Sprintf("sum of invalid ids: %d", invalidIds)
	return solution, nil
}

func sumInvalidIds(rangeMin int64, rangeMax int64, validationFunc func(id int64) bool, verboseOutput bool) int64 {

	var invalidIdSum int64 = 0
	if verboseOutput {
		fmt.Printf("min: %d; max: %d\n", rangeMin, rangeMax)
	}
	for id := rangeMin; id <= rangeMax; id++ {

		if validationFunc(id) {
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
func isValidId2(id int64) bool {

	idString := fmt.Sprint(id)

	// grab slices from start until the midway point, to detect repeating text
	// (any larger than half the length means there are no repeating sequences, thus a valid Id )
	maxSliceLength := len(idString) / 2
	for sliceLength := 1; sliceLength <= maxSliceLength; sliceLength++ {

		// has a repeating slice, invalid
		if hasRepeatingSlice(idString, sliceLength) {
			return false
		}
	}

	// still here, assume valid
	return true
}
func hasRepeatingSlice(idString string, sliceLength int) bool {

	sliceCount := float64(len(idString)) / float64(sliceLength)
	if sliceCount != math.Floor(sliceCount) {
		// uneven slices, can't be repeating
		return false
	}

	if sliceCount < 2 {
		// nothing to compare
		return false
	}
	firstSlice := idString[:sliceLength]
	// compare slices
	for sliceIndex := 1; sliceIndex < int(sliceCount); sliceIndex++ {

		nextSliceIndex := sliceLength * sliceIndex
		nextSlice := idString[nextSliceIndex : nextSliceIndex+sliceLength]

		// interruption of repeating sequence
		if firstSlice != nextSlice {
			return false
		}
	}
	// still here, slices are repeating
	return true
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
