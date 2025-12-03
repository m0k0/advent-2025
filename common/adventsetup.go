package common

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AdventSetup struct {
	Day           int32  `yaml:"day"`
	Variant       string `yaml:"variant"`
	Input         string `yaml:"input"`
	VerboseOutput bool   `yaml:"verboseOutput"`
}

func (advent *AdventSetup) ReadFromYamlFile(path string) error {

	configFile, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("i/o error: \n%w", err)
	}

	err = yaml.Unmarshal(configFile, advent)
	if err != nil {
		return fmt.Errorf("yaml parse error: \n%w", err)
	}

	return nil
}

func (advent *AdventSetup) GetDataDirPath() (string, error) {
	if advent.Day < 1 || advent.Day > 25 {
		return "", errors.New("invalid day value")
	}
	dataDir := "day"
	if advent.Day < 10 {
		dataDir += "0"
	}
	dataDir += fmt.Sprintf("%d", advent.Day)

	return dataDir, nil
}

func (advent *AdventSetup) GetFilePath(fileType string) (string, error) {
	dataDir, err := advent.GetDataDirPath()
	if err != nil {
		return "", fmt.Errorf("failed to determine data directory: %w", err)
	}
	path := fmt.Sprintf(`%s/%s.%s.txt`, dataDir, advent.Input, fileType)

	return path, nil
}

func (advent *AdventSetup) OpenInput() (*os.File, error) {

	path, err := advent.GetFilePath("input")
	if err != nil {
		return nil, fmt.Errorf("failed to get input file path: %w", err)
	}

	inputFile, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to read input '%s': \n%w", advent.Input, err)
	}

	return inputFile, nil
}

func (advent *AdventSetup) ReadSolution() (string, error) {
	path, err := advent.GetFilePath("solution")
	if err != nil {
		return "", fmt.Errorf("failed to get solution file path: %w", err)
	}

	solution, err := os.ReadFile(path)
	if !os.IsExist(err) {
		return "", nil
	}
	if err != nil {
		// other I/O error
		return "", fmt.Errorf("failed to read solution for '%s': \n%w", advent.Input, err)
	}

	return string(solution), nil
}
