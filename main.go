package main

import (
	"errors"
	"fmt"
	"m0k0/advent-2025/day01"
	"os"

	"gopkg.in/yaml.v3"
)

type AdventConfig struct {
	Day     int32  `yaml:"day"`
	Variant string `yaml:"variant"`
	Input   string `yaml:"input"`
}

func main() {
	fmt.Println("❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️")
	fmt.Println("🎄  Advent of Code 2025  🎄")
	fmt.Println("☃️✨🎁~~~~~~~~~~~🎁✨☃️")

	config, err := readConfig("config.yaml")
	if err != nil {
		fmt.Print(fmt.Errorf("failed to read config: \n%w", err))
		os.Exit(1)
	}

	fmt.Printf("Working on a solution for day %d/%s, using '%s' input...\n\n",
		config.Day,
		config.Variant,
		config.Input)

	solution, err := solve(config)
	if err != nil {
		fmt.Print(fmt.Errorf("failed to solve: \n%w", err))
		os.Exit(1)
	}

	fmt.Print(solution)
}

func readConfig(path string) (*AdventConfig, error) {

	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("i/o error: \n%w", err)
	}

	config := &AdventConfig{}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return nil, fmt.Errorf("yaml parse error: \n%w", err)
	}

	return config, nil
}

func solve(config *AdventConfig) (string, error) {
	switch config.Day {
	case 1:
		return day01.Solve(config.Variant, config.Input)
	default:
		return "", errors.New("no solution availble for this day")
	}
}
