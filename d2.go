package aoc2023

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day2(input string, part int) (int, error) {
	games := strings.Split(input, "\n")

	var possible, minimum int
	for _, game := range games {
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

func day2alt(input string, part int) (int, error) {
	r := strings.NewReader(input)
	s := bufio.NewScanner(r)

	var a1 int
	var a2 int

	for s.Scan() {
		l := s.Text()

		// skip 'Game '
		i := 5

		// game index
		var gi int
		for isDigit(rune(l[i])) {
			gi *= 10
			gi += int(l[i] - '0')
			i++
		}

		// skip ': '
		i += 2

		// max cube counts
		var ci, r, g, b int
		for i < len(l) {
			ru := rune(l[i])
			switch {
			case isDigit(ru):
				ci *= 10
				ci += int(ru - '0')
				i++
			case ru == 'r':
				if ci > r {
					r = ci
				}
				ci = 0
				i += 3
			case ru == 'g':
				if ci > g {
					g = ci
				}
				i += 5
				ci = 0
			case ru == 'b':
				if ci > b {
					b = ci
				}
				i += 4
				ci = 0
			default:
				i++
			}
		}

		if r <= 12 && g <= 13 && b <= 14 {
			a1 += gi
		}

		a2 += r * b * g
	}

	if part == 1 {
		return a1, nil
	}

	return a2, nil
}

func isDigit(r rune) bool {
	return r-'0' >= 0 && r-'0' <= 9
}
