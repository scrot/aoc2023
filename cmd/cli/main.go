package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/scrot/aoc2023"
)

var cmds = map[string]*flag.FlagSet{
	"run":       flag.NewFlagSet("run", flag.ExitOnError),
	"profile":   flag.NewFlagSet("profile", flag.ExitOnError),
	"test":      flag.NewFlagSet("test", flag.ExitOnError),
	"benchmark": flag.NewFlagSet("benchmark", flag.ExitOnError),
}

func main() {
	// parse universal flags
	var day int
	for _, c := range cmds {
		c.IntVar(&day, "day", 1, "puzzle day")
	}

	// parse specific flags
	var version, part int
	cmds["run"].IntVar(&version, "version", 1, "version of puzzle")
	cmds["run"].IntVar(&part, "part", 1, "part of the puzzle (1 or 2)")
	cmds["profile"].IntVar(&version, "version", 1, "version of puzzle")
	cmds["profile"].IntVar(&part, "part", 1, "part of the puzzle (1 or 2)")

	sub, ok := cmds[os.Args[1]]
	if !ok {
		log.Fatalf("invalid subcommand %s", os.Args[1])
	}
	sub.Parse(os.Args[2:])

	switch os.Args[1] {
	case "run":
		path := fmt.Sprintf("./day%d/input.txt", day)
		input, _ := os.ReadFile(path)
		var answer int
		answer, err := aoc2023.Days[day][version-1].Solve(input, part)
		if err != nil {
			fmt.Printf("V%d: error %q\n", version, err)
			break
		}
		fmt.Printf("V%d: Answer day %d part %d is %d\n", version, day, part, answer)
	case "benchmark":
		path := fmt.Sprintf("./day%d", day)
		c := exec.Command("go", "test", "-run='^$'", "-bench=.", path)
		c.Stdout = os.Stdout
		if err := c.Run(); err != nil {
			log.Fatalf("%s: %s", c.String(), err)
		}
	case "test":
		path := fmt.Sprintf("./day%d/.", day)
		c := exec.Command("go", "test", path)
		c.Stdout = os.Stdout
		if err := c.Run(); err != nil {
			log.Fatalf("%s: %s", c.String(), err)
		}
	case "profile":
		path := fmt.Sprintf("./day%d", day)
		out := fmt.Sprintf("%s/day%d-part%d-version%d-cpu.pprof", path, day, part, version)
		benchmark := fmt.Sprintf("-bench=BenchmarkDay%dPart%dV%d", day, part, version)
		c := exec.Command("go", "test", "-run='^$'", benchmark, "-cpuprofile", out, path)
		if err := c.Run(); err != nil {
			log.Fatalf("%s: %s", c.String(), err)
		}

		c = exec.Command("go", "tool", "pprof", "-web", "./aoc2023", out)
		c.Stdout = os.Stdout
		if err := c.Run(); err != nil {
			log.Fatalf("%s: %s", c.String(), err)
		}
	}
}
