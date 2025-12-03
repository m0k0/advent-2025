package main

import (
	"errors"
	"fmt"
	"m0k0/advent-2025/common"
	"m0k0/advent-2025/day01"
	"m0k0/advent-2025/day02"
	"m0k0/advent-2025/day03"
	"os"
)

func main() {
	fmt.Println("❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️")
	fmt.Println("🎄  Advent of Code 2025  🎄")
	fmt.Println("☃️✨🎁~~~~~~~~~~~🎁✨☃️")

	solver := &common.AdventSetup{}
	err := solver.ReadFromYamlFile("config.yaml")

	if err != nil {
		fmt.Print(fmt.Errorf("failed to read config: \n%w", err))
		os.Exit(1)
	}

	fmt.Printf("Working on a solution for day %d/%s, using '%s' input...\n\n",
		solver.Day,
		solver.Variant,
		solver.Input)

	solution, err := solve(solver)
	if err != nil {
		fmt.Print(fmt.Errorf("failed to solve: \n%w", err))
		os.Exit(1)
	}

	fmt.Print(solution)
}

func solve(advent *common.AdventSetup) (string, error) {
	switch advent.Day {
	case 1:
		return day01.Solve(advent)
	case 2:
		return day02.Solve(advent)
	case 3:
		return day03.Solve(advent)
	default:
		return "", errors.New("no solution availble for this day")
	}
}
