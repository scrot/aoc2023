# Advent of code 2023

<!--toc:start-->
- [Advent of code 2023](#advent-of-code-2023)
  - [Run solver](#run-solver)
  - [Run tests](#run-tests)
  - [Run benchmarks](#run-benchmarks)
<!--toc:end-->

## Run solver

Build and run the program and use `day` to select which puzzle to solve `version` to select the version of the solver (if available) and `part` to solve the first part or second part of the puzzle.

```shell
> go build . && ./aoc2023 run -day 1 -version 1 -part 1
```

## Run tests

```shell
❯ go build . && ./aoc2023 test -day 1
```

## Run benchmarks

```shell
❯ go build . && ./aoc2023 benchmark -day 1
```

## Run Profiler

Profile generates a `pprof` cpu profile and generates a graph that is opened in the browser. it requires `graphviz` to work (e.g. install it using `brew install graphviz`).

The profiler uses benchmarks that are formatted `BenchmarkDay1Part1V1` to generate a cpu profile.

```shell
> go build . && ./aoc2023 profile -day 1 -version 1 -part 1
```
