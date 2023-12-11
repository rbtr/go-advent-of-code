package common

import (
	"log"
	"time"
)

const (
	sampleFileName = "sample"
	inputFileName  = "input"
)

type Puzzle struct {
	Sample Data
	Input  Data
}

func Load() (p *Puzzle, err error) {
	p = &Puzzle{}
	if p.Sample, err = Ingest(sampleFileName); err != nil {
		return nil, err
	}
	if p.Input, err = Ingest(inputFileName); err != nil {
		return nil, err
	}
	return p, nil
}

type Solver func(d Data) (string, error)

func Solve(p *Puzzle, solvers ...Solver) {
	for i, solver := range solvers {
		start := time.Now()
		sample, err := solver(p.Sample)
		if err != nil {
			log.Fatal(err)
		}
		duration := time.Since(start)
		log.Printf("%d.sample\t%s\t(%s)", i+1, sample, duration.String())
		start = time.Now()
		input, err := solver(p.Input)
		if err != nil {
			log.Fatal(err)
		}
		duration = time.Since(start)
		log.Printf("%d.input\t%s\t(%s)", i+1, input, duration.String())
	}
}
