package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	common "github.com/rbtr/go-advent-of-code/2021"
)

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := one(puzzle); err != nil {
		log.Fatal(err)
	}
	if err := two(puzzle); err != nil {
		log.Fatal(err)
	}
	if err := twoButMath(puzzle); err != nil {
		log.Fatal(err)
	}
}

func one(puzzle *common.Puzzle) error {
	sc := bufio.NewScanner(bytes.NewReader(puzzle.Input))
	sc.Split(bufio.ScanLines)
	sc.Scan()
	splits := strings.Split(sc.Text(), ",")
	arr := make([]int, len(splits))
	var err error
	var max int
	for i := range splits {
		if arr[i], err = strconv.Atoi(splits[i]); err != nil {
			return err
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	crabs := make([][]int, len(arr))
	for i := range arr {
		pos := make([]int, max+1)
		for j := range pos {
			pos[j] = int(math.Abs(float64(j - arr[i])))
		}
		crabs[i] = pos
	}

	sums := make([]int, max+1)
	for i := range crabs {
		for j := range crabs[i] {
			sums[j] += crabs[i][j]
		}
	}

	min := sums[0]
	idx := 0
	for i := range sums {
		if sums[i] <= min {
			idx = i
			min = sums[i]
		}
	}

	fmt.Printf("max: %d, num: %d, %d: %d\n", max, len(crabs), idx, min)

	return nil
}

func two(puzzle *common.Puzzle) error {
	sc := bufio.NewScanner(bytes.NewReader(puzzle.Input))
	sc.Split(bufio.ScanLines)
	sc.Scan()
	splits := strings.Split(sc.Text(), ",")
	arr := make([]int, len(splits))
	var err error
	var max int
	for i := range splits {
		if arr[i], err = strconv.Atoi(splits[i]); err != nil {
			return err
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	crabs := make([][]int, len(arr))
	for i := range arr {
		pos := make([]int, max+1)
		for j := range pos {
			del := int(math.Abs(float64(j - arr[i])))
			sum := 0
			for i := del; i > 0; i-- {
				sum += i
			}
			pos[j] = sum
		}
		crabs[i] = pos
	}

	sums := make([]int, max+1)
	for i := range crabs {
		for j := range crabs[i] {
			sums[j] += crabs[i][j]
		}
	}

	min := sums[0]
	idx := 0
	for i := range sums {
		if sums[i] <= min {
			idx = i
			min = sums[i]
		}
	}

	fmt.Printf("max: %d, num: %d, %d: %d\n", max, len(crabs), idx, min)
	return nil
}

func twoButMath(puzzle *common.Puzzle) error {
	sc := bufio.NewScanner(bytes.NewReader(puzzle.Input))
	sc.Split(bufio.ScanLines)
	sc.Scan()
	splits := strings.Split(sc.Text(), ",")
	arr := make([]int, len(splits))
	var err error
	var max int
	for i := range splits {
		if arr[i], err = strconv.Atoi(splits[i]); err != nil {
			return err
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	crabs := make([][]int, len(arr))
	for i := range arr {
		pos := make([]int, max+1)
		for j := range pos {
			n := int(math.Abs(float64(j - arr[i])))
			pos[j] = (int(math.Pow(float64(n), 2)) + n) / 2
		}
		crabs[i] = pos
	}

	sums := make([]int, max+1)
	for i := range crabs {
		for j := range crabs[i] {
			sums[j] += crabs[i][j]
		}
	}

	min := sums[0]
	idx := 0
	for i := range sums {
		if sums[i] <= min {
			idx = i
			min = sums[i]
		}
	}

	fmt.Printf("max: %d, num: %d, %d: %d\n", max, len(crabs), idx, min)
	return nil
}
