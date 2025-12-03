package day01_test

import (
	"fmt"
	"m0k0/advent-2025/common"
	"m0k0/advent-2025/day01"
	"testing"
)

func TestDay01(t *testing.T) {

	var tests = []*common.AdventSetup{
		{
			Day:     1,
			Variant: "part1",
			Input:   "test",
		},
	}

	for i := range tests {
		advent := tests[i]
		name := fmt.Sprintf("variant_%s_input_%s", advent.Variant, advent.Input)
		expectedSolution, err := advent.ReadSolution()
		if err != nil {
			t.Error(err)
			t.Fatal("error reading solution")
		}

		t.Run(name, func(t *testing.T) {

			solution, err := day01.Solve(advent)
			if err != nil {
				t.Error(err)
				t.Fatal("error while solving")
			}
			if solution != expectedSolution {
				t.Fatalf("invalid solution; expected '%s', produced '%s'", expectedSolution, solution)
			}
		})
	}
}
