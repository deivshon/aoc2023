package main

import (
	"main/src/pkg/day1"
	"main/src/pkg/day2"
	"main/src/pkg/day3"
)

type Solution struct {
	Day         int
	SolveFirst  func() (string, error)
	SolveSecond func() (string, error)
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
}
