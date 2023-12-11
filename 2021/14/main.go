package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	common "github.com/rbtr/aoc2021"
)

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	common.Solve(puzzle, one, two)
}

func one(puzzle []byte) (string, error) {
	lines, err := common.SplitLines(puzzle)
	if err != nil {
		return "", err
	}
	polymer := lines[0][0]
	subs := map[string]string{}
	for i := 2; i < len(lines[0]); i++ {
		instrs := strings.Split(lines[0][i], " -> ")
		subs[instrs[0]] = instrs[1]
	}
	iterations := 10
	for i := 0; i < iterations; i++ {
		substitutions := map[int]string{}
		for s := 0; s <= len(polymer)-2; s++ {
			pair := polymer[s : s+2]
			substitutions[s+1] = subs[pair]
		}
		idx := []int{}
		for j := range substitutions {
			idx = append(idx, j)
		}
		sort.Ints(idx)
		for j := idx[len(idx)-1]; j >= 0; j-- {
			if repl, ok := substitutions[j]; ok {
				polymer = polymer[:j] + repl + polymer[j:]
			}
		}
	}
	counts := map[string]int{}
	for i := range polymer {
		counts[string(polymer[i])] = counts[string(polymer[i])] + 1
	}
	max := 0
	min := -1
	for key := range counts {
		if min < 0 {
			min = counts[key]
		}
		if counts[key] < min {
			min = counts[key]
		}
		if counts[key] > max {
			max = counts[key]
		}
	}
	return fmt.Sprintf("%d", max-min), nil
}

func two(puzzle []byte) (string, error) {
	lines, err := common.SplitLines(puzzle)
	if err != nil {
		return "", err
	}
	polymer := lines[0][0]
	subs := map[string]string{}
	for i := 2; i < len(lines[0]); i++ {
		instrs := strings.Split(lines[0][i], " -> ")
		subs[instrs[0]] = instrs[1]
	}
	counts := map[string]int{}
	for s := range polymer {
		counts[string(polymer[s])] = counts[string(polymer[s])] + 1
	}
	pairs := map[string]int{}
	for s := 0; s <= len(polymer)-2; s++ {
		pairs[polymer[s:s+2]] = pairs[polymer[s:s+2]] + 1
	}
	iterations := 40
	for i := 0; i < iterations; i++ {
		newPairs := map[string]int{}
		for pair, count := range pairs {
			if sub, ok := subs[pair]; ok {
				counts[sub] = counts[sub] + count
				newPair0 := string(pair[0]) + sub
				newPair1 := sub + string(pair[1])
				newPairs[newPair0] = newPairs[newPair0] + count
				newPairs[newPair1] = newPairs[newPair1] + count
			} else {
				newPairs[pair] = newPairs[pair] + count
			}
		}
		pairs = newPairs
	}
	max := 0
	min := -1
	for key := range counts {
		if min < 0 {
			min = counts[key]
		}
		if counts[key] < min {
			min = counts[key]
		}
		if counts[key] > max {
			max = counts[key]
		}
	}
	return fmt.Sprintf("%d", max-min), nil
}
