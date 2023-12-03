package main

import (
	"fmt"
	"main/src/pkg/utils"
	"time"
)

func main() {
	for _, s := range Solutions {
		start := time.Now()
		firstSolution, err := s.SolveFirst()
		end := time.Now()

		if err != nil {
			utils.Failure(fmt.Sprintf("first solution of day %v returned an erorr: %v", s.Day, err))
		}
		elapsedFirst := end.Sub(start)

		start = time.Now()
		secondSolution, err := s.SolveSecond()
		end = time.Now()

		if err != nil {
			utils.Failure(fmt.Sprintf("second solution of day %v returned an erorr: %v", s.Day, err))
		}
		elapsedSecond := end.Sub(start)

		fmt.Printf("day %02d: %8v - %-10v | %8v - %-10v\n", s.Day, firstSolution, elapsedFirst, secondSolution, elapsedSecond)
	}
}
