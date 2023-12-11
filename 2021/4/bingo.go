package main

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type board struct {
	hasWon     bool
	rowTouches []int
	colTouches []int
	sum        int
}

type fieldMapper struct {
	row, col int
	b        *board
}

type bingo struct {
	boards int
	draws  []int
	field  map[int][]fieldMapper
}

func newBingo(b []byte) (*bingo, error) {
	bingo := &bingo{
		field: map[int][]fieldMapper{},
	}

	sc := bufio.NewScanner(bytes.NewReader(b))
	sc.Split(bufio.ScanLines)
	sc.Scan()

	drawsStr := strings.Split(sc.Text(), ",")
	bingo.draws = make([]int, len(drawsStr))
	for i := range drawsStr {
		draw, err := strconv.Atoi(drawsStr[i])
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert string to int")
		}
		bingo.draws[i] = draw
	}

	i := 0
	var bd *board
	var sw *bufio.Scanner
	for sc.Scan() {
		if sc.Text() == "" {
			continue
		}
		if i%5 == 0 {
			bd = &board{
				rowTouches: []int{0, 0, 0, 0, 0},
				colTouches: []int{0, 0, 0, 0, 0},
				sum:        0,
			}
			bingo.boards++
		}
		sw = bufio.NewScanner(bytes.NewReader(sc.Bytes()))
		sw.Split(bufio.ScanWords)
		j := 0
		for sw.Scan() {
			s := sw.Text()
			val, err := strconv.Atoi(s)
			if err != nil {
				return nil, errors.Wrap(err, "failed to convert string to int")
			}
			bd.sum += val
			bingo.field[val] = append(bingo.field[val], fieldMapper{row: i % 5, col: j, b: bd})
			j++
		}
		i++
	}
	return bingo, nil
}
