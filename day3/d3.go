package day3

import "fmt"

func Solve(input []byte, part, version int) (int, error) {
	switch version {
	case 1:
		return SolveV1(input, part)
	case 2:
		// return SolveV2(input, part)
		return 0, nil
	default:
		return 0, fmt.Errorf("invalid version %d", version)
	}
}
