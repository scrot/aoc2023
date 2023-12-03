package day1

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func day1Profile(input []byte, part, version int) {
	file, _ := os.Create("day1-part1.pprof")
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	i, _ := Solve(input, part, version)
	fmt.Printf("Day 1 solution of part 1: %d", i)
}
