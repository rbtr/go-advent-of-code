package puzzle

import (
	"fmt"
	"time"
)

func Run(d Data, solvers ...Solver) error {
	var start time.Time
	for i := range solvers {
		start = time.Now()
		res, err := solvers[i](d)
		if err != nil {
			return err
		}
		end := time.Since(start)
		fmt.Printf("%d: %s (%s)\n", i+1, res, end.String())
	}
	return nil
}
