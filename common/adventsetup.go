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

func (advent *AdventSetup) OpenInput() (*os.File, error) {

	if advent.Day < 1 || advent.Day > 25 {
		return nil, errors.New("invalid day value")
	}
	dataDir := "day"
	if advent.Day < 10 {
		dataDir += "0"
	}
	dataDir += fmt.Sprintf("%d", advent.Day)

	path := fmt.Sprintf(`%s/%s.input.txt`, dataDir, advent.Input)

	inputFile, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to read input '%s': \n%w", advent.Input, err)
	}

	return inputFile, nil
}
