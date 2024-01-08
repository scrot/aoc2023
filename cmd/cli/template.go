package main

import (
	"fmt"
	"io"
	"text/template"
	"unicode"
)

var funcMap = template.FuncMap{
	"title": func(s string) string {
		return fmt.Sprintf("%c%s", unicode.ToUpper(rune(s[0])), s[1:])
	},
}

func writeTemplate(templ string, w io.WriteCloser, dir string) error {
	t, err := template.New("template").Funcs(funcMap).Parse(templ)
	if err != nil {
		return err
	}

	err = t.Execute(w, dir)
	if err != nil {
		return err
	}

	return nil
}

const puzzleImpl = `package {{.}}

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	return 0, nil
}
`

const puzzleTest = `package {{.}}_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/{{.}}"
)

//go:embed input.txt
var input []byte

const example = ""

func Test{{title .}}(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Example", 1, {{.}}.V1{}, []byte(example), 0},
	}

	for _, c := range cs {
		got, _ := c.version.Solve(c.input, c.part)
		if got != c.want {
			t.Errorf("want %d got %d", c.want, got)
		}
	}
}

func benchmark{{title .}}(b *testing.B, s aoc2023.Solver, part int) {
	for i := 0; i < b.N; i++ {
		s.Solve(input, part)
	}
}
func Benchmark{{title .}}Part1V1(b *testing.B) { benchmark{{title .}}(b, {{.}}.V1{}, 1) }
`

const puzzleInput = ``
