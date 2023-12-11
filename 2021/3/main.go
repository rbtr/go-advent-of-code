package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	common "github.com/rbtr/aoc2021"
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
}

func collapseMatr(rows [][]int) []int {
	sum := make([]int, len(rows[0]))
	for i := range rows {
		for j := range rows[i] {
			sum[j] += rows[i][j]
		}
	}
	return sum
}

func strToInts(s string) ([]int, error) {
	ints := make([]int, len(s))
	var err error
	for i := range s {
		ints[i], err = strconv.Atoi(string(s[i]))
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func strMatrToIntMatr(s []string) ([][]int, error) {
	res := make([][]int, len(s))
	var err error
	for i := range s {
		res[i], err = strToInts(s[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func setDigits(ints []int, hi, lo, count int) []int {
	out := make([]int, len(ints))
	for i := range ints {
		if count%2 == 0 && ints[i] == count/2 {
			out[i] = hi
			continue
		}
		if ints[i] > count/2 {
			out[i] = hi
			continue
		}
		out[i] = lo
	}
	return out
}

func arrToString(ints []int) string {
	out := make([]string, len(ints))
	for i := range ints {
		out[i] = strconv.Itoa(ints[i])
	}
	return strings.Join(out, "")
}

func parseBinaryString(s ...string) ([]int, error) {
	res := make([]int, len(s)+1)
	for i := range s {
		ii, err := strconv.ParseInt(s[i], 2, 64)
		if err != nil {
			return nil, err
		}
		res[i] = int(ii)
	}
	res[len(res)-1] = res[len(res)-2] * res[len(res)-3]
	return res, nil
}

func one(puzzle *common.Puzzle) error {
	lines, err := common.SplitLines(puzzle.Sample, puzzle.Input)
	if err != nil {
		return err
	}
	intMatr, err := strMatrToIntMatr(lines[0])
	if err != nil {
		return err
	}
	sums := collapseMatr(intMatr)
	decs, err := parseBinaryString(arrToString(setDigits(sums, 1, 0, len(intMatr))), arrToString(setDigits(sums, 0, 1, len(intMatr))))
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", decs)

	intMatr, err = strMatrToIntMatr(lines[1])
	if err != nil {
		return err
	}
	sums = collapseMatr(intMatr)
	decs, err = parseBinaryString(arrToString(setDigits(sums, 1, 0, len(intMatr))), arrToString(setDigits(sums, 0, 1, len(intMatr))))
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", decs)
	return nil
}

func dropRows(ints [][]int, idx, hi, lo int) []int {
	if len(ints) == 1 {
		return ints[0]
	}
	sums := collapseMatr(ints)
	comp := setDigits(sums, hi, lo, len(ints))
	sub := make([][]int, 0)
	for i := range ints {
		if ints[i][idx] == comp[idx] {
			sub = append(sub, ints[i])
		}
	}
	return dropRows(sub, idx+1, hi, lo)
}

func two(puzzle *common.Puzzle) error {
	lines, err := common.SplitLines(puzzle.Sample, puzzle.Input)
	if err != nil {
		return err
	}
	intMatr, err := strMatrToIntMatr(lines[0])
	if err != nil {
		return err
	}
	dec, err := parseBinaryString(arrToString(dropRows(intMatr, 0, 1, 0)), arrToString(dropRows(intMatr, 0, 0, 1)))
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", dec)
	intMatr, err = strMatrToIntMatr(lines[1])
	if err != nil {
		return err
	}
	dec, err = parseBinaryString(arrToString(dropRows(intMatr, 0, 1, 0)), arrToString(dropRows(intMatr, 0, 0, 1)))
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", dec)
	return nil
}
