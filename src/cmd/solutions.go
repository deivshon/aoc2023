package main

import (
	"main/src/pkg/day1"
	"main/src/pkg/day2"
	"main/src/pkg/day3"
	"main/src/pkg/day4"
	"main/src/pkg/day5"
	"main/src/pkg/day6"
)

const SolutionNotComputed = "not computed"

type Solution struct {
	Day                     int
	SolveFirst              func() (string, error)
	SolveSecond             func() (string, error)
	InefficientResultFirst  *InefficientResult
	InefficientResultSecond *InefficientResult
}

type InefficientResult struct {
	BenchmarkedTime string
}

var Solutions = []Solution{
	{
		Day:         1,
		SolveFirst:  day1.SolveFirst,
		SolveSecond: day1.SolveSecond,
	},
	{
		Day:         2,
		SolveFirst:  day2.SolveFirst,
		SolveSecond: day2.SolveSecond,
	},
	{
		Day:         3,
		SolveFirst:  day3.SolveFirst,
		SolveSecond: day3.SolveSecond,
	},
	{
		Day:         4,
		SolveFirst:  day4.SolveFirst,
		SolveSecond: day4.SolveSecond,
	},
	{
		Day:         5,
		SolveFirst:  day5.SolveFirst,
		SolveSecond: day5.SolveSecond,
		InefficientResultSecond: &InefficientResult{
			BenchmarkedTime: "1m20s",
		},
	},
	{
		Day:         6,
		SolveFirst:  day6.SolveFirst,
		SolveSecond: day6.SolveSecond,
	},
}
