package puzzle

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	rootCtx   context.Context
	filenames = []string{"sample", "input"}
)

type Raw []byte

type Solver func(d Raw) (string, error)

type Puzzle struct {
	inputs []Raw
}

func Load() (p *Puzzle, err error) {
	p = &Puzzle{inputs: []Raw{nil, nil}}
	for i := range filenames {
		if p.inputs[i], err = os.ReadFile(filenames[i]); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func (p *Puzzle) Run(solvers ...Solver) error {
	var start time.Time
	for i := range solvers {
		for j := range p.inputs {
			start = time.Now()
			res, err := solvers[i](p.inputs[j])
			if err != nil {
				return err
			}
			end := time.Since(start)
			fmt.Printf("%d-%s: %s (%s)\n", i+1, filenames[j], res, end.String())
		}
	}
	return nil
}

// init() is executed before main() whenever this package is imported
// to do pre-run setup of things like exit signal handling and building
// the root context.
func init() {
	var cancel context.CancelFunc
	rootCtx, cancel = context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Wait until receiving a signal.
		sig := <-sigCh
		log.Printf("caught exit signal %v, exiting\n", sig)
		cancel()
		log.Printf("exiting")
		os.Exit(1)
	}()
}
