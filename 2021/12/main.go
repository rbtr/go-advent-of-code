package main

import (
	"fmt"
	"log"
	"strings"

	common "github.com/rbtr/aoc2021"
)

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := one(puzzle.Sample); err != nil {
		log.Fatal(err)
	}
	if err := one(puzzle.Input); err != nil {
		log.Fatal(err)
	}
	if err := two(puzzle.Sample); err != nil {
		log.Fatal(err)
	}
	if err := two(puzzle.Input); err != nil {
		log.Fatal(err)
	}
}

func bigCave(s string) bool {
	return s == strings.ToUpper(s)
}

func copyMap(m map[string]int) map[string]int {
	out := map[string]int{}
	for k, v := range m {
		out[k] = v
	}
	return out
}

func traverse(current string, graph map[string]map[string]interface{}, visits map[string]int, path string, paths map[string]interface{}) {
	path += current
	visits[current]++
	if current == "end" {
		paths[path] = nil
	}
	for next := range graph[current] {
		times := visits[next]
		if next == "start" {
			continue
		}
		if !bigCave(next) && times > 0 {
			continue
		}
		traverse(next, graph, copyMap(visits), path, paths)
	}
}

func one(puzzle []byte) error {
	lines, err := common.SplitLines(puzzle)
	if err != nil {
		return err
	}

	connections := map[string]map[string]interface{}{}
	for _, line := range lines[0] {
		arc := strings.Split(line, "-")
		dests, ok := connections[arc[0]]
		if !ok {
			dests = map[string]interface{}{}
		}
		dests[arc[1]] = nil
		connections[arc[0]] = dests

		src, ok := connections[arc[1]]
		if !ok {
			src = map[string]interface{}{}
		}
		src[arc[0]] = nil
		connections[arc[1]] = src
	}
	visits := map[string]int{}
	paths := map[string]interface{}{}
	traverse("start", connections, visits, "", paths)
	fmt.Println(len(paths))
	return nil
}

func two(puzzle []byte) error {
	lines, err := common.SplitLines(puzzle)
	if err != nil {
		return err
	}
	connections := map[string]map[string]interface{}{}
	for _, line := range lines[0] {
		arc := strings.Split(line, "-")
		dests, ok := connections[arc[0]]
		if !ok {
			dests = map[string]interface{}{}
		}
		dests[arc[1]] = nil
		connections[arc[0]] = dests

		src, ok := connections[arc[1]]
		if !ok {
			src = map[string]interface{}{}
		}
		src[arc[0]] = nil
		connections[arc[1]] = src
	}
	paths := map[string]interface{}{}
	for connection := range connections {
		if bigCave(connection) || connection == "start" || connection == "end" {
			continue
		}
		visits := map[string]int{
			connection: -1,
		}
		traverse("start", connections, visits, "", paths)
	}
	fmt.Println(len(paths))
	return nil
}
