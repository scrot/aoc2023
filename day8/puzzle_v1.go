package day8

import (
	"bufio"
	"bytes"
)

type V1 struct{}

type node struct {
	name        string
	left, right string
}

func newNode(l string) node {
	return node{l[0:3], l[7:10], l[12:15]}
}

type cycle struct {
	index                  int
	match, prevmatch, diff int
	current                string
}

func (V1) Solve(input []byte, part int) (int, error) {
	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	s.Scan()
	instructions := s.Text()
	s.Scan()

	nodes := make(map[string]node)
	for s.Scan() {
		n := newNode(s.Text())
		nodes[n.name] = n
	}

	var (
		steps int
		cs    = []node{nodes["AAA"]}
	)

	if part == 2 {
		cs = nodesWithEnding('A', nodes)
	}

	cycles := make(map[string]cycle)
	for steps < 100000 {
		for _, i := range instructions {
			if part == 1 && cs[0].name == "ZZZ" {
				return steps, nil
			}

			if part == 2 && allEndsWith('Z', cs) {
				return steps, nil
			}

			temp := cs
			cs = []node{}
			for j, c := range temp {
				if endsWith('Z', c) {

					// first occurance of ..Z for node
					if _, ok := cycles[c.name]; !ok {
						cycles[c.name] = cycle{j, steps, 0, steps, c.name}
					}

					// second occurance of ..Z for node, record difference
					tmp := cycles[c.name]
					cycles[c.name] = cycle{j, steps, tmp.match, steps - tmp.match, c.name}

					// all nodes have made at least one cycle
					// multiply all diffs between cycles for first sync
					var res []int
					for _, cycle := range cycles {
						if cycle.diff != 0 {
							res = append(res, cycle.diff)
						}
					}

					if len(res) == 6 {
						return lcm(res[0], res[1], res[2:]...), nil
					}

				}
				if i == 'L' {
					cs = append(cs, nodes[c.left])
				} else {
					cs = append(cs, nodes[c.right])
				}
			}
			steps++

		}
	}
	return 0, nil
}

func allEndsWith(r rune, ns []node) bool {
	for _, n := range ns {
		if !endsWith(r, n) {
			return false
		}
	}
	return true
}

func endsWith(r rune, n node) bool {
	return rune(n.name[2]) == r
}

func nodesWithEnding(r rune, ns map[string]node) []node {
	var matches []node
	for k := range ns {
		if endsWith(r, ns[k]) {
			matches = append(matches, ns[k])
		}
	}
	return matches
}

// greatest common divisor (gcd) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (lcm) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
