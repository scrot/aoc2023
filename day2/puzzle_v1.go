package day2

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type V1 struct{}

func (_ V1) Solve(input []byte, part int) (int, error) {
	if part > 2 {
		return 0, errors.New("invalid part")
	}

	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	var possible, minimum int
	for s.Scan() {
		game := s.Text()

		i := index(game)
		sets := sets(game)
		if validCount(sets, 12, 13, 14) {
			possible += i
		}

		r, g, b := minCubes(sets)
		minimum += r * g * b

	}

	if part == 2 {
		return minimum, nil
	}

	return possible, nil
}

func index(l string) int {
	re := regexp.MustCompile(`Game (\d+):.*`)
	matches := re.FindStringSubmatch(l)
	i, _ := strconv.Atoi(matches[1])
	return i
}

func sets(l string) []string {
	re := regexp.MustCompile(`Game \d+: (.*)`)
	matches := re.FindStringSubmatch(l)
	sets := strings.Split(matches[1], ";")
	return sets
}

func validCount(sets []string, red, green, blue int) bool {
	for _, set := range sets {
		cubes := strings.Split(set, ", ")
		for _, c := range cubes {
			var (
				count  int
				colour string
			)

			fmt.Sscanf(c, "%d %s", &count, &colour)

			switch {
			case colour == "red" && count > red:
				return false
			case colour == "green" && count > green:
				return false
			case colour == "blue" && count > blue:
				return false
			}
		}
	}
	return true
}

func minCubes(sets []string) (r int, g int, b int) {
	for _, set := range sets {
		cubes := strings.Split(set, ", ")
		for _, c := range cubes {
			var (
				count  int
				colour string
			)

			fmt.Sscanf(c, "%d %s", &count, &colour)

			switch {
			case colour == "red" && count > r:
				r = count
			case colour == "green" && count > g:
				g = count
			case colour == "blue" && count > b:
				b = count
			}
		}
	}
	return r, g, b
}
