package day03

import (
	"bufio"
	"fmt"
	"m0k0/advent-2025/common"
	"math"
	"strconv"
	"sync"
)

func Solve(advent *common.AdventSetup) (string, error) {

	inputFile, err := advent.OpenInput()
	if err != nil {
		return "", fmt.Errorf("day 2 error: \n%w", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var numberOfBatteries int64 = 2
	if advent.Variant == "part2" {
		numberOfBatteries = 12
	}
	var totalJoltage int64 = 0
	bankNumber := 0

	var tasks sync.WaitGroup

	for scanner.Scan() {
		bankNumber++
		bankText := scanner.Text()

		//fmt.Print(bankText)

		tasks.Add(1)
		go func() {
			maxJoltage, err := getMaxJoltage(bankText, numberOfBatteries, 0, advent.VerboseOutput)
			if err != nil {
				fmt.Print(fmt.Errorf("error on bank %d, '%s': %w", bankNumber, bankText, err))
			}
			totalJoltage += maxJoltage
			tasks.Done()
		}()

	}

	tasks.Wait()

	solution := fmt.Sprintf("total output joltage: %d", totalJoltage)

	return solution, nil
}

func getMaxJoltage(batteryBank string, batteryCount int64, startingJoltage int64, verboseOutput bool) (int64, error) {

	var maxJoltage int64 = 0
	var joltage int64 = 0
	var reservedBatteries [2]int64

	if verboseOutput {
		fmt.Printf("bank: %s\n", batteryBank)
	}

	for i1 := 0; i1 < len(batteryBank); i1++ {

		isEndOfBank := i1 >= len(batteryBank)-int(batteryCount)-1
		if startingJoltage == 0 {
			isEndOfBank = i1 > len(batteryBank)-int(batteryCount)
			if isEndOfBank {
				break
			}
		}

		battery1, err := strconv.ParseInt(string(batteryBank[i1]), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse battery '%c': %w", batteryBank[i1], err)
		}

		joltage = battery1 * int64(math.Pow(10, float64(batteryCount-1)))

		if joltage < maxJoltage {
			continue
		}

		if batteryCount > 1 {
			subJoltage, err := getMaxJoltage(batteryBank[i1+1:], batteryCount-1, joltage, verboseOutput)
			if err != nil {
				return 0, err
			}
			joltage += subJoltage
		}

		if startingJoltage > 0 {
			potentialJoltage := joltage * 10
			if !isEndOfBank && startingJoltage > 0 && potentialJoltage > startingJoltage {
				return -startingJoltage, nil
			}
		}

		if joltage > maxJoltage {
			maxJoltage = joltage
		}
	}

	if verboseOutput {
		tabs := ""
		for i := 0; i < int(batteryCount); i++ {
			tabs += "\t"
		}
		fmt.Printf("%smax joltage: %d\n", tabs, maxJoltage)
	}

	return maxJoltage, nil
	/*
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
	}*/

	maxJoltageText := fmt.Sprintf("%d%d", reservedBatteries[0], reservedBatteries[1])

	maxJoltage, err := strconv.ParseInt(maxJoltageText, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse joltage '%s': %w", maxJoltageText, err)
	}

	return maxJoltage, nil
}
