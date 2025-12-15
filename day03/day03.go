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

	logger := common.Logger{
		Verbose: advent.VerboseOutput,
	}

	for scanner.Scan() {
		bankNumber++
		bankText := scanner.Text()

		tasks.Add(1)
		go func() {

			maxJoltage, err := getMaxJoltageReverse(bankText, numberOfBatteries, logger)
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

func getMaxJoltageReverse(batteryBank string, batteryCount int64, logger common.Logger) (int64, error) {

	logger.PrintVerboseF("bank: %s\n", batteryBank)

	var startingIndex = len(batteryBank) - int(batteryCount)
	var largestBattery int64 = 0
	var largestBatteryIndex int = -1

	for i := startingIndex; i >= 0; i-- {
		battery, err := strconv.ParseInt(string(batteryBank[i]), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse battery '%c': %w", batteryBank[i], err)
		}

		if battery >= largestBattery {
			largestBatteryIndex = i
			largestBattery = battery
		}
	}

	joltage := largestBattery * int64(math.Pow(10, float64(batteryCount-1)))

	logger.PrintVerboseFD("joltage: %d\n", int(batteryCount), joltage)

	if batteryCount > 1 {
		subJoltage, err := getMaxJoltageReverse(batteryBank[largestBatteryIndex+1:], batteryCount-1, logger)
		if err != nil {
			return 0, err
		}
		joltage += subJoltage
	}

	logger.PrintVerboseFD("max joltage: %d\n", int(batteryCount), joltage)

	return joltage, nil

}
