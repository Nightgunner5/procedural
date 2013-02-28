package main

import (
	"flag"
	"fmt"
	"github.com/Nightgunner5/procedural/data"
	"github.com/Nightgunner5/procedural/export"
	"os"
	"runtime/pprof"
	"time"
)

var (
	seed    = flag.Int64("seed", time.Now().UnixNano(), "a number that the world will be based on, using horribly complex mathmatical functions. defaults to the number of nanoseconds since midnight (GMT) on 1970-01-01.")
	cpuprof = flag.String("cpuprof", "", "if non-empty, the name of a file that will contain a CPU profile.")
)

func main() {
	flag.Parse()

	if *cpuprof != "" {
		f, err := os.Create(*cpuprof)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var w data.World
	w.Seed = *seed
	w.Init()

	f, err := os.Create(fmt.Sprintf("seed_%d.html", *seed))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = export.Export(f, &w)
	if err != nil {
		panic(err)
	}
}
