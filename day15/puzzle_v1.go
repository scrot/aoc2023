package day15

import "fmt"

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	var sum int

	var hash int
	for _, b := range input {
		switch b {
		case ',':
			fmt.Printf("hash: %d sum: %d\n", hash, sum)
			sum += hash
			hash = 0
		case '\n', '\r', '\036':
			// ignore newlines
		default:
			hash += int(b)
			hash *= 17
			hash %= 256
		}
	}
	sum += hash
	fmt.Printf("hash: %d sum: %d\n", hash, sum)

	return sum, nil
}
