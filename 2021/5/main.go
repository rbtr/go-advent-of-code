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
	// if err := one(puzzle); err != nil {
	// 	log.Fatal(err)
	// }
	if err := two(puzzle); err != nil {
		log.Fatal(err)
	}
}

type coord struct {
	x, y int
}

func parse(b []byte) ([][]coord, error) {
	res := [][]coord{}
	sl := bufio.NewScanner(bytes.NewReader(b))
	sl.Split(bufio.ScanLines)
	for sl.Scan() {
		seg := []coord{}
		sw := bufio.NewScanner(bytes.NewBuffer(sl.Bytes()))
		sw.Split(bufio.ScanWords)
		for sw.Scan() {
			if sw.Text() == "->" {
				continue
			}
			xy := strings.Split(sw.Text(), ",")
			point := coord{}
			var err error
			if point.x, err = strconv.Atoi(xy[0]); err != nil {
				return nil, err
			}
			if point.y, err = strconv.Atoi(xy[1]); err != nil {
				return nil, err
			}
			seg = append(seg, point)
		}
		res = append(res, seg)
	}
	return res, nil
}

func dropDiags(segments [][]coord) [][]coord {
	res := [][]coord{}
	for _, segment := range segments {
		if segment[0].x == segment[1].x || segment[0].y == segment[1].y {
			res = append(res, segment)
		}
	}
	return res
}

func one(puzzle *common.Puzzle) error {
	segments, err := parse(puzzle.Input)
	if err != nil {
		return err
	}
	orthos := dropDiags(segments)

	pts := map[coord]int{}
	for _, segment := range orthos {
		x := segment[0].x
		y := segment[0].y
		for x != segment[1].x || y != segment[1].y {
			c := coord{x: x, y: y}
			pts[c] = pts[c] + 1
			if x != segment[1].x {
				x += ((segment[1].x - segment[0].x) / int(math.Abs(float64(segment[1].x-segment[0].x))))
			}
			if y != segment[1].y {
				y += ((segment[1].y - segment[0].y) / int(math.Abs(float64(segment[1].y-segment[0].y))))
			}
		}
		c := coord{x: x, y: y}
		pts[c] = pts[c] + 1
	}
	mult := 0
	for _, v := range pts {
		if v > 1 {
			mult++
		}
	}

	fmt.Printf("%v\n%v\n%+v\n%d\n", segments, orthos, pts, mult)
	return nil
}

func two(puzzle *common.Puzzle) error {
	segments, err := parse(puzzle.Input)
	if err != nil {
		return err
	}
	// orthos := dropDiags(segments)

	pts := map[coord]int{}
	for _, segment := range segments {
		x := segment[0].x
		y := segment[0].y
		for x != segment[1].x || y != segment[1].y {
			c := coord{x: x, y: y}
			pts[c] = pts[c] + 1
			if x != segment[1].x {
				x += ((segment[1].x - segment[0].x) / int(math.Abs(float64(segment[1].x-segment[0].x))))
			}
			if y != segment[1].y {
				y += ((segment[1].y - segment[0].y) / int(math.Abs(float64(segment[1].y-segment[0].y))))
			}
		}
		c := coord{x: x, y: y}
		pts[c] = pts[c] + 1
	}
	mult := 0
	for _, v := range pts {
		if v > 1 {
			mult++
		}
	}

	fmt.Printf("%v\n%v\n%d\n", segments, pts, mult)
	return nil
}
