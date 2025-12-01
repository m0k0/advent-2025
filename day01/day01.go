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
	var rotation int64 = 50
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

		switch line[0] {
		case 'L':
			rotationMovement *= -1
		case 'R':
			rotationMovement *= 1
		default:
			return "", fmt.Errorf("invalid instruction on line %d: '%s'", lineNumber, line)
		}

		rotation += rotationMovement
		if rotation < DIAL_MIN {
			rotation = DIAL_MAX + (rotation % DIAL_MAX)
		}
		if rotation >= DIAL_MAX {
			rotation = DIAL_MIN + (rotation % DIAL_MAX)
		}

		if rotation == 0 {
			zeroCount++
		}
	}

	solution := fmt.Sprintf("the password is: %d", zeroCount)

	return solution, nil
}
