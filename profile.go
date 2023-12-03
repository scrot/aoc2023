package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func Profiler(input []byte, day, part, version int) string {
	path := fmt.Sprintf("day%d-part%d-version%d.prof", day, part, version)
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("invalid path %s: %s", path, err)
	}

	if err := pprof.StartCPUProfile(file); err != nil {
		log.Fatalf("profile of %s: %s", path, err)
	}
	defer pprof.StopCPUProfile()

	if _, err := days[day](input, part, version); err != nil {
		log.Fatalf("profiling %s: %s", path, err)
	}

	return path
}
