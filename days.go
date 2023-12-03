package main

import (
	"github.com/scrot/aoc2023/day1"
	"github.com/scrot/aoc2023/day2"
)

var days = map[int]func([]byte, int, int) (int, error){
	1: day1.Solve,
	2: day2.Solve,
}
