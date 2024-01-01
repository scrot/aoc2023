package day15

import (
	"slices"
)

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	var (
		sum     int
		hashmap = make(map[hash][]lens, 256)

		seq   []byte
		label []byte
		flen  int
		op    byte
	)

	for _, b := range input {
		switch b {
		case ',':
			h := hash(0)
			h.Write(seq)
			sum += int(h)
			execute(op, lens{label, flen}, hashmap)
			seq, label = []byte{}, []byte{}
		case '\n', '\r', '\036':
			// ignore newlines
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			flen = int(b - '0')
			seq = append(seq, b)
		case '=', '-':
			op = b
			if b == '-' {
				flen = 0
			}
			seq = append(seq, b)
		default:
			label = append(label, b)
			seq = append(seq, b)
		}
	}

	h := hash(0)
	h.Write(seq)
	sum += int(h)
	execute(op, lens{label, flen}, hashmap)

	if part == 2 {
		var power int
		for i, boxes := range hashmap {
			for j, box := range boxes {
				power += (int(i) + 1) * (j + 1) * box.flen
			}
		}
		return power, nil
	}

	return sum, nil
}

type hash int

func (h *hash) Write(data []byte) (int, error) {
	for _, b := range data {
		*h += hash(b)
		*h *= 17
		*h %= 256
	}
	return len(data), nil
}

type lens struct {
	label []byte
	flen  int
}

func (l lens) hash() hash {
	h := hash(0)
	h.Write(l.label[:])
	return h
}

func execute(op byte, l lens, hm map[hash][]lens) {
	var (
		hi = l.hash()
		ls = hm[hi]
	)

	if op == '=' && len(ls) == 0 {
		hm[l.hash()] = []lens{l}
		return
	}

	for i := range ls {
		if slices.Equal(l.label, ls[i].label) {
			switch op {
			case '=':
				hm[hi][i] = l
			case '-':
				hm[hi] = append(ls[:i], ls[i+1:]...)
			}
			return
		}
	}

	if op == '=' {
		hm[hi] = append(hm[hi], l)
	}
}
