package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func openInput(inputName string) (*os.File, error) {

	path := fmt.Sprintf(`day01/%s.input.txt`, inputName)

	inputFile, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to read input '%s': \n%w", inputName, err)
	}

	return inputFile, nil
}

func Solve(variant string, input string) (string, error) {

	inputFile, err := openInput(input)
	if err != nil {
		return "", fmt.Errorf("day 1 error: \n%w", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	const DIAL_MIN = 0
	const DIAL_MAX = 100
	var zeroCount int64 = 0
	var clickCount int64 = 0
	var currentRotation int64 = 50
	var lineNumber int64 = 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		if len(line) < 2 {
			return "", fmt.Errorf("invalid input on line %d: '%s'", lineNumber, line)
		}

		rotationMovement, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			return "", fmt.Errorf("invalid rotation number on line %d: '%s'", lineNumber, line)
		}

		var rotationVector int64 = 0
		switch line[0] {
		case 'L':
			rotationVector = -1
		case 'R':
			rotationVector = 1
		default:
			return "", fmt.Errorf("invalid instruction on line %d: '%s'", lineNumber, line)
		}

		for rotationMovement != 0 {
			currentCycleMovement := (rotationMovement % DIAL_MAX)
			if currentCycleMovement == 0 {
				currentCycleMovement = DIAL_MAX
			}
			rotationMovement -= currentCycleMovement
			newRotation := currentRotation + currentCycleMovement*rotationVector

			// don't click if already on zero, unless doing a full turn
			if (currentRotation != 0 || currentCycleMovement == DIAL_MAX) &&
				(newRotation <= DIAL_MIN || newRotation >= DIAL_MAX) {
				clickCount++
			}

			if newRotation <= DIAL_MIN {
				newRotation = DIAL_MAX + (newRotation % DIAL_MAX)
			} else if newRotation >= DIAL_MAX {
				newRotation = DIAL_MIN + (newRotation % DIAL_MAX)
			}

			if newRotation == DIAL_MAX {
				newRotation = 0
				zeroCount++
			}

			currentRotation = newRotation
		}

	}

	solution := ""
	switch variant {
	case "part1":
		solution = fmt.Sprintf("the password is: %d", zeroCount)
	case "part2":
		solution = fmt.Sprintf("the password is: %d", clickCount)
	}

	return solution, nil
}
