package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"

	common "github.com/rbtr/go-advent-of-code/2021"
)

func main() {
	err := os.Mkdir("test", 0o777)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		fmt.Println("err")
	}
	fmt.Println("okay")
	// puzzle, err := common.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// common.Solve(puzzle, one, two)
}

type card struct {
	id   int
	have []int
	want []int
	val  int
}

func parseCard(line string) card {
	c := card{}
	split := strings.Split(line, ":")
	c.id, _ = strconv.Atoi(strings.Split(split[0], " ")[1])
	field := strings.Split(split[1], "|")
	for _, win := range strings.Split(field[0], " ") {
		s := strings.TrimSpace(win)
		if s == "" {
			continue
		}
		i, _ := strconv.Atoi(s)
		c.want = append(c.want, i)
	}
	for _, num := range strings.Split(field[1], " ") {
		s := strings.TrimSpace(num)
		if s == "" {
			continue
		}
		i, _ := strconv.Atoi(s)
		c.have = append(c.have, i)
	}
	return c
}

func one(p common.Data) (string, error) {
	cards := []card{}
	for _, l := range p.AsLines() {
		cards = append(cards, parseCard(l))
	}
	total := 0
	for i := range cards {
		hits := 0
		for _, potential := range cards[i].have {
			for _, winner := range cards[i].want {
				if potential == winner {
					hits++
				}
			}
		}
		if hits == 0 {
			continue
		}
		val := 1 << (hits - 1)
		// fmt.Printf("card %d val %d\n", cards[i].id, val)
		total += val
	}
	return fmt.Sprintf("%d", total), nil
}

func two(p common.Data) (string, error) {
	cards := []card{}
	for _, l := range p.AsLines() {
		cards = append(cards, parseCard(l))
	}
	for i := range cards {
		for _, potential := range cards[i].have {
			for _, winner := range cards[i].want {
				if potential == winner {
					cards[i].val++
				}
			}
		}
	}
	fmt.Println(cards)
	frames := [][]card{cards}
	for o := 0; o < len(frames); o++ {
		prev := frames[o]
		next := []card{}
		for _, c := range prev {
			for i := 0; i < c.val; i++ {
				idx := c.id + i
				fmt.Printf("card %d adds %d (%d - %d)\n", c.id, c.val, idx, cards[idx].id)
				next = append(next, cards[idx])
			}
			frames = append(frames, next)
		}
	}
	total := 0
	for i := range frames {
		total += len(frames[i])
	}
	return fmt.Sprintf("%d", total), nil
}
