package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	common "github.com/rbtr/go-advent-of-code/2021"
)

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	common.Solve(puzzle, one, two)
}

type color string

const (
	red   color = "red"
	green color = "green"
	blue  color = "blue"
)

type game struct {
	id    int
	shown []map[color]int
}

func parseGame(line string) game {
	game := game{}
	g := strings.Split(line, ":")
	gameID := strings.TrimPrefix(g[0], "Game ")
	game.id, _ = strconv.Atoi(gameID)
	sets := strings.Split(g[1], ";")
	for _, set := range sets {
		shown := map[color]int{}
		diceCounts := strings.Split(set, ",")
		for _, dice := range diceCounts {
			dice = strings.TrimSpace(dice)
			tokens := strings.Split(dice, " ")
			count, _ := strconv.Atoi(tokens[0])
			shown[color(tokens[1])] = count
		}
		game.shown = append(game.shown, shown)
	}
	return game
}

func one(p common.Data) (string, error) {
	games := []game{}
	for _, l := range p.AsLines() {
		games = append(games, parseGame(l))
	}
	bound := map[color]int{
		red:   12,
		green: 13,
		blue:  14,
	}
	sum := 0
	for _, g := range games {
		possible := true
		for _, shown := range g.shown {
			if !possible {
				continue
			}
			for color, count := range shown {
				if !possible {
					continue
				}
				if count > bound[color] {
					possible = false
				}
			}
		}
		if !possible {
			continue
		}
		fmt.Printf("game %d %t\n", g.id, possible)
		sum += g.id
	}

	return fmt.Sprintf("%d", sum), nil
}

func two(p common.Data) (string, error) {
	games := []game{}
	for _, l := range p.AsLines() {
		games = append(games, parseGame(l))
	}
	sum := 0
	for _, g := range games {
		min := map[color]int{
			red:   0,
			green: 0,
			blue:  0,
		}
		for _, shown := range g.shown {
			for color, count := range shown {
				if count > min[color] {
					min[color] = count
				}
			}
		}
		sum += min[red] * min[green] * min[blue]
	}

	return fmt.Sprintf("%d", sum), nil
}
