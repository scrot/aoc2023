package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/scrot/aoc2023/day1"
	"github.com/scrot/aoc2023/day2"
)

var (
	day     int
	part    int
	version int
)

var cmds = map[string]*flag.FlagSet{
	"run":    flag.NewFlagSet("run", flag.ExitOnError),
	"profle": flag.NewFlagSet("profile", flag.ExitOnError),
	"test":   flag.NewFlagSet("test", flag.ExitOnError),
}

func main() {
	// parse universal flags
	for _, c := range cmds {
		c.IntVar(&day, "day", 1, "puzzle day")
		c.IntVar(&version, "version", 1, "version of puzzle")
		c.IntVar(&part, "part", 1, "part of the puzzle (1 or 2)")
	}

	// parse specific flags
	sub, ok := cmds[os.Args[1]]
	if !ok {
		log.Fatalf("invalid subcommand %s", os.Args[1])
	}
	sub.Parse(os.Args[2:])

	switch os.Args[1] {
	case "run":
		run()
	case "profile":
	case "test":
	}
}

func run() {
	path := fmt.Sprintf("./day%d/input.txt", day)
	input, _ := os.ReadFile(path)

	var answer int

	switch day {
	case 1:
		answer, _ = day1.Solve(input, part, version)
	case 2:
		answer, _ = day2.Solve(input, part, version)
	}

	fmt.Printf("Answer day %d part %d: %d\n", day, part, answer)
}
