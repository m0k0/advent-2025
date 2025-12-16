package day04

import (
	"bufio"
	"fmt"
	"m0k0/advent-2025/common"
)

const PAPER rune = '@'
const SPACE rune = '.'

const SEARCH_RADIUS = 1
const SEARCH_GRID_SIZE = SEARCH_RADIUS*2 + 1
const REQUIRED_ROWS = SEARCH_RADIUS * 2
const MAX_ADJACENT_ROLLS = 3

func Solve(advent *common.AdventSetup) (string, error) {

	inputFile, err := advent.OpenInput()
	if err != nil {
		return "", fmt.Errorf("day 4 error: \n%w", err)
	}
	defer inputFile.Close()

	logger := common.Logger{
		Verbose: advent.VerboseOutput,
	}

	scanner := bufio.NewScanner(inputFile)

	grid := common.Grid[rune]{}
	rowIx := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid.SetValues(rowIx, []rune(line))
		rowIx++
	}

	gridArray := grid.Slice(0, 0, grid.Width, grid.Height)

	for y := range gridArray {
		logger.PrintVerboseF("%s\n", string(gridArray[y]))
	}

	return "", nil
	rowQueue := common.Queue[string]{}
	rowIndex := 0
	totalAccessibleRolls := 0
	for {

		hasData := scanner.Scan()
		if hasData {
			rowQueue.Push(scanner.Text())
		}

		if rowQueue.Length < REQUIRED_ROWS {
			if !hasData {
				break
			}
			// not enough data to assess adjacency, read more
			continue
		}

		// create search grid
		var searchGrid [SEARCH_GRID_SIZE]string

		i := 0
		if rowIndex < SEARCH_RADIUS {
			i = rowIndex + SEARCH_RADIUS
		}
		logger.PrintVerbose("Search grid: \n")
		for item := range rowQueue.Items() {
			searchGrid[i] = item
			i++

			if advent.VerboseOutput {
				marker := "]"
				if i == len(searchGrid)/2+1 {
					marker = ">"
				}
				logger.PrintVerboseF("%s %s\n", marker, item)
			}
		}
		// assess grid
		logger.PrintVerbose("Accessible rolls of paper: \n")
		accessibleRolls := countAccessibleRollsOfPaper(searchGrid[:], MAX_ADJACENT_ROLLS, logger)
		logger.PrintVerboseF("\nFound %d accessible rolls \n\n", accessibleRolls)
		totalAccessibleRolls += accessibleRolls

		if rowIndex >= SEARCH_RADIUS {
			// storing more than we need
			rowQueue.Pop()
		}
		rowIndex++
	}
	solution := fmt.Sprintf("Total accessible rolls of paper: %d\n", totalAccessibleRolls)
	return solution, nil
}

func countAccessibleRollsOfPaper(searchGrid []string, maxRolls int, logger common.Logger) int {

	searchRadius := len(searchGrid) / 2
	rowIndex := searchRadius
	rowToAnalyse := searchGrid[rowIndex]

	accessibleRolls := 0
	logger.PrintVerbose("  ")
	for charIndex := 0; charIndex < len(rowToAnalyse); charIndex++ {

		value := rune(rowToAnalyse[charIndex])

		if value != PAPER {
			logger.PrintVerbose(".")
			continue
		}

		// found paper, check if accessible
		adjacentRollsOfPaper := 0
		for y := -searchRadius; y <= searchRadius; y++ {
			if rowIndex+y < 0 || rowIndex+y >= len(rowToAnalyse) {
				continue //out of bounds
			}
			if searchGrid[rowIndex+y] == "" {
				continue // empty row, skip
			}

			for x := -searchRadius; x <= searchRadius; x++ {
				if charIndex+x < 0 || charIndex+x >= len(rowToAnalyse) {
					continue //out of bounds
				}

				if x == 0 && y == 0 {
					continue // skip current
				}

				adjacentValue := rune(searchGrid[rowIndex+y][charIndex+x])

				if adjacentValue != PAPER {
					continue // not paper
				}

				adjacentRollsOfPaper++
				if adjacentRollsOfPaper > maxRolls {
					break // too many adjacent rolls
				}
			}
			if adjacentRollsOfPaper > maxRolls {
				break // too many adjacent rolls
			}
		}

		if adjacentRollsOfPaper > maxRolls {
			logger.PrintVerbose(".")
			continue // too many adjacent rolls
		}

		logger.PrintVerbose("x")
		accessibleRolls++ // still here, so accessible
	}
	return accessibleRolls
}
