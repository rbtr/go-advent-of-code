package main

import (
	"fmt"
	"log"
	"strings"

	common "github.com/rbtr/aoc2021"
)

// var strategy = map[string]string{
// 	"A": "Y", // rock, paper
// 	"B": "X", // paper, rock
// 	"C": "Z", // scissors, scissors
// }

var scores = map[string]int{
	"X": 1, // rock
	"Y": 2, // paper
	"Z": 3, // scissors
}

var opts = []string{"X", "Y", "Z"}

var results = []int{0, 3, 6}

func main() {
	two()
}

func one() {
	p, err := common.Load()
	if err != nil {
		log.Println(err)
	}
	m, _ := p.Input.AsStringMatrix()

	for _, round := range m {
		switch round[0] {
		case "A":
			round[0] = "X"
		case "B":
			round[0] = "Y"
		case "C":
			round[0] = "Z"
		}
	}
	fmt.Println(m)

	total := 0
	for _, round := range m {
		fmt.Println(round)
		// tie
		if round[0] == round[2] {
			total += (3 + scores[round[2]])
			fmt.Println("tie", total)
			continue
		}
		// index of
		idx0 := scores[round[0]] - 1
		idx1 := scores[round[2]] - 1
		fmt.Println(idx0, idx1)
		// win
		if idx1 == (idx0+1)%3 {
			total += (6 + scores[round[2]])
			fmt.Println("win", total)
			continue
		}
		// else loss
		total += (0 + scores[round[2]])
		fmt.Println("loss", total)
	}
	fmt.Println(total)
}

var play = "ABC"

func two() {
	p, err := common.Load()
	if err != nil {
		log.Println(err)
	}
	m, _ := p.Input.AsStringMatrix()
	fmt.Println(m)

	total := 0
	for _, round := range m {
		idx := strings.Index(play, round[0])
		choose := idx
		switch round[2] {
		case "X": // lose
			choose -= 1
			if choose < 0 {
				choose = 2
			}
			total += choose + 1
		case "Y": // draw
			total += choose + 1 + 3
		case "Z": // win
			choose += 1
			choose = choose % 3
			total += choose + 1 + 6
		}
		fmt.Println(string(play[choose]))

	}
	fmt.Println(total)
}
