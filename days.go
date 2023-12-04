package aoc2023

import (
	"github.com/scrot/aoc2023/day1"
	"github.com/scrot/aoc2023/day2"
	"github.com/scrot/aoc2023/day3"
	"github.com/scrot/aoc2023/day4"
)

type Solver interface {
	Solve([]byte, int) (int, error)
}

var Days = map[int][]Solver{
	1: {day1.V1{}, day1.V2{}},
	2: {day2.V1{}, day2.V1{}},
	3: {day3.V1{}},
	4: {day4.V1{}},
}
