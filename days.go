package aoc2023

import (
	"github.com/scrot/aoc2023/day1"
	"github.com/scrot/aoc2023/day10"
	"github.com/scrot/aoc2023/day11"
	"github.com/scrot/aoc2023/day12"
	"github.com/scrot/aoc2023/day13"
	"github.com/scrot/aoc2023/day2"
	"github.com/scrot/aoc2023/day3"
	"github.com/scrot/aoc2023/day4"
	"github.com/scrot/aoc2023/day5"
	"github.com/scrot/aoc2023/day6"
	"github.com/scrot/aoc2023/day7"
	"github.com/scrot/aoc2023/day8"
	"github.com/scrot/aoc2023/day9"
)

type Solver interface {
	Solve([]byte, int) (int, error)
}

var Days = map[int][]Solver{
	1:  {day1.V1{}, day1.V2{}},
	2:  {day2.V1{}, day2.V1{}},
	3:  {day3.V1{}},
	4:  {day4.V1{}},
	5:  {day5.V1{}},
	6:  {day6.V1{}, day6.V2{}},
	7:  {day7.V1{}},
	8:  {day8.V1{}},
	9:  {day9.V1{}},
	10: {day10.V1{}},
	11: {day11.V1{}},
	12: {day12.V1{}},
	13: {day13.V1{}},
}
