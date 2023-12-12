package day9

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
	"strings"
)

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	// read all sequences
	var seqs [][]int
	for s.Scan() {
		numbers := strings.Split(s.Text(), " ")
		var seq []int
		for _, n := range numbers {
			x, _ := strconv.Atoi(n)
			seq = append(seq, x)
		}
		seqs = append(seqs, seq)
	}

	var res int
	for _, seq := range seqs {
		switch part {
		case 1:
			res += nextIn(seq, false)
		case 2:
			res += nextIn(seq, true)
		default:
			return 0, errors.New("invalid part")
		}
	}

	return res, nil
}

func nextIn(seq []int, backwards bool) int {
	var diffs []int

	for i := 0; i < len(seq)-1; i += 1 {
		diff := seq[i+1] - seq[i]
		diffs = append(diffs, diff)
	}

	// tail: zero differences
	if allZero(diffs) {
		if backwards {
			return seq[0]
		}
		return seq[len(seq)-1]
	}

	if backwards {
		return seq[0] - nextIn(diffs, backwards)
	}
	return seq[len(seq)-1] + nextIn(diffs, backwards)
}

func allZero(ns []int) bool {
	zeros := true
	for _, n := range ns {
		if n != 0 {
			zeros = false
		}
	}
	return zeros
}
