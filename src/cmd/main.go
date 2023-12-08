package main

import (
	"flag"
	"fmt"
	"main/src/pkg/utils"
	"time"
)

func main() {
	var computeInefficient = flag.Bool("c", false, "Compute inefficient solutions")
	flag.Parse()

	if !*computeInefficient {
		fmt.Printf("Benchmarked times for solutions not computed on the fly are indicated with * (use -c flag to compute all solutions)\n")
		fmt.Printf("Benchmark system specifications\n\t+ CPU\tRyzen 5 3600\n\t+ RAM\t32GiB\n\n")
	} else {
		fmt.Printf("All solutions, including inefficient ones, are going to be computed on the fly\n\n")
	}

	for _, s := range Solutions {
		var firstSolution string
		var elapsedFirst string
		if s.InefficientResultFirst != nil && !*computeInefficient {
			firstSolution = SolutionNotComputed
			elapsedFirst = fmt.Sprintf("~%s*", s.InefficientResultFirst.BenchmarkedTime)
		} else {
			start := time.Now()
			computedFirstSolution, err := s.SolveFirst()
			end := time.Now()

			if err != nil {
				utils.Failure(fmt.Sprintf("first solution of day %v returned an erorr: %v", s.Day, err))
			}

			firstSolution = computedFirstSolution
			elapsedFirst = fmt.Sprint(end.Sub(start))
		}

		var secondSolution string
		var elapsedSecond string
		if s.InefficientResultSecond != nil && !*computeInefficient {
			secondSolution = SolutionNotComputed
			elapsedSecond = fmt.Sprintf("~%s*", s.InefficientResultSecond.BenchmarkedTime)
		} else {
			start := time.Now()
			secondComputedSolution, err := s.SolveSecond()
			end := time.Now()

			if err != nil {
				utils.Failure(fmt.Sprintf("second solution of day %v returned an error: %v", s.Day, err))
			}

			secondSolution = secondComputedSolution
			elapsedSecond = fmt.Sprint(end.Sub(start))
		}

		fmt.Printf("day %02d: %-15v%-10v | %-15v%-10v\n", s.Day, firstSolution, elapsedFirst, secondSolution, elapsedSecond)
	}
}
