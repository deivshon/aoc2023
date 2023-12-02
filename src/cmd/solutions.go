package main

import "main/src/pkg/day1"

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
}
