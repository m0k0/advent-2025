package common

import "fmt"

type Logger struct {
	Verbose    bool
	IndentText string
}

func (logger *Logger) PrintVerbose(message string) {

	if !logger.Verbose {
		return
	}
	fmt.Print(message)
}

func (logger *Logger) PrintVerboseF(template string, params ...any) {

	if !logger.Verbose {
		return
	}
	fmt.Printf(template, params...)
}
func (logger *Logger) PrintVerboseFD(template string, depth int, params ...any) {

	if !logger.Verbose {
		return
	}
	indentation := ""

	if len(logger.IndentText) == 0 {
		logger.IndentText = " "
	}

	for i := 0; i < depth; i++ {
		indentation += logger.IndentText
	}

	template = indentation + template

	fmt.Printf(template, params...)
}
