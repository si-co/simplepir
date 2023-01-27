package pir

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

const LOGQ = uint64(32)
const SEC_PARAM = uint64(1 << 10)

// Benchmark SimplePIR performance.
func TestSimplePirSingle(t *testing.T) {
	f, err := os.Create("simple-cpu.out")
	if err != nil {
		panic("Error creating file")
	}

	d := uint64(2048)

	// initialize experiment
	experiment := &Experiment{Results: make(map[int][]*Chunk, 0)}

	for _, log_N := range []uint64{2, 12, 22} {
		N := uint64(1 << log_N)

		pir := SimplePIR{}
		p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

		i := uint64(0) // index to query
		if i >= p.l*p.m {
			panic("Index out of dimensions")
		}

		DB := MakeRandomDB(N, d, &p)
		var tputs []float64

		nRepeat := 30
		results := make([]*Chunk, nRepeat)
		var tput float64
		for j := 0; j < nRepeat; j++ {
			results[j], tput, _, _, _ = RunFakePIR(&pir, DB, p, []uint64{i}, f, false)
			tputs = append(tputs, tput)
		}
		experiment.Results[1<<(int(log_N)+11)] = results
	}

	res, err := json.Marshal(experiment)
	if err != nil {
		panic(err)
	}
	fileName := "simplePIR" + ".json"
	if err = ioutil.WriteFile(fileName, res, 0644); err != nil {
		panic(err)
	}
}

// Benchmark DoublePIR performance.
func TestDoublePirSingle(t *testing.T) {
	f, err := os.Create("double-cpu.out")
	if err != nil {
		panic("Error creating file")
	}

	d := uint64(2048)

	// initialize experiment
	experiment := &Experiment{Results: make(map[int][]*Chunk, 0)}

	for _, log_N := range []uint64{2, 12, 22} {
		N := uint64(1 << log_N)

		pir := DoublePIR{}
		p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

		i := uint64(0) // index to query
		if i >= p.l*p.m {
			panic("Index out of dimensions")
		}

		DB := MakeRandomDB(N, d, &p)
		var tputs []float64

		nRepeat := 30
		results := make([]*Chunk, nRepeat)
		var tput float64
		for j := 0; j < nRepeat; j++ {
			results[j], tput, _, _, _ = RunFakePIR(&pir, DB, p, []uint64{i}, f, false)
			tputs = append(tputs, tput)
		}

		experiment.Results[1<<(int(log_N)+11)] = results
	}

	res, err := json.Marshal(experiment)
	if err != nil {
		panic(err)
	}
	fileName := "doublePIR" + ".json"
	if err = ioutil.WriteFile(fileName, res, 0644); err != nil {
		panic(err)
	}
}
