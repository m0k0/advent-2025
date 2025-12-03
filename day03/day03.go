package day03

import (
	"bufio"
	"fmt"
	"m0k0/advent-2025/common"
	"strconv"
)

func Solve(advent *common.AdventSetup) (string, error) {

	inputFile, err := advent.OpenInput()
	if err != nil {
		return "", fmt.Errorf("day 2 error: \n%w", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var totalJoltage int64 = 0
	bankNumber := 0
	for scanner.Scan() {
		bankNumber++
		bankText := scanner.Text()

		maxJoltage, err := getMaxJoltage(bankText, advent.VerboseOutput)
		if err != nil {
			return "", fmt.Errorf("error on bank %d, '%s': %w", bankNumber, bankText, err)
		}
		totalJoltage += maxJoltage
	}

	solution := fmt.Sprintf("total output joltage: %d", totalJoltage)

	return solution, nil
}

func getMaxJoltage(batteryBank string, verboseOutput bool) (int64, error) {

	var maxJoltage int64 = 0
	var reservedBatteries [2]int64

	if verboseOutput {
		fmt.Printf("bank: %s\n", batteryBank)
	}
	for i1 := 0; i1 < len(batteryBank)-1; i1++ {

		battery1, err := strconv.ParseInt(string(batteryBank[i1]), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse battery '%c': %w", batteryBank[i1], err)
		}
		if battery1 > reservedBatteries[0] {
			reservedBatteries[0] = battery1
			reservedBatteries[1] = 0
		}

		for i2 := i1 + 1; i2 < len(batteryBank); i2++ {

			endOfBank := i2 >= len(batteryBank)-1

			if verboseOutput && endOfBank {
				fmt.Print("\tend of bank\n")
			}

			battery2, err := strconv.ParseInt(string(batteryBank[i2]), 10, 64)
			if err != nil {
				return 0, fmt.Errorf("failed to parse battery '%c': %w", batteryBank[i2], err)
			}
			if battery2 > battery1 && !endOfBank {
				break
			}
			if battery2 > reservedBatteries[1] {
				reservedBatteries[1] = battery2
			}
		}
	}

	maxJoltageText := fmt.Sprintf("%d%d", reservedBatteries[0], reservedBatteries[1])

	if verboseOutput {
		fmt.Printf("\tmax joltage: %s\n", maxJoltageText)
	}

	maxJoltage, err := strconv.ParseInt(maxJoltageText, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse joltage '%s': %w", maxJoltageText, err)
	}

	return maxJoltage, nil
}
