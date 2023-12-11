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
	common.Solve(puzzle, one, two)
}

var hexBin = map[rune][]int{
	'0': {0, 0, 0, 0},
	'1': {0, 0, 0, 1},
	'2': {0, 0, 1, 0},
	'3': {0, 0, 1, 1},
	'4': {0, 1, 0, 0},
	'5': {0, 1, 0, 1},
	'6': {0, 1, 1, 0},
	'7': {0, 1, 1, 1},
	'8': {1, 0, 0, 0},
	'9': {1, 0, 0, 1},
	'A': {1, 0, 1, 0},
	'B': {1, 0, 1, 1},
	'C': {1, 1, 0, 0},
	'D': {1, 1, 0, 1},
	'E': {1, 1, 1, 0},
	'F': {1, 1, 1, 1},
}

func hexToBin(s string) []int {
	out := []int{}
	for _, r := range s {
		out = append(out, hexBin[r]...)
	}
	return out
}

func bitsToInt(in ...int) int {
	s := ""
	for _, i := range in {
		s += strconv.Itoa(i)
	}
	out, _ := strconv.ParseInt(s, 2, 64)
	return int(out)
}

type packet struct {
	ver int
	typ int
	val int
	sub []packet
}

func parse(in []int, num int) (int, []packet) {
	counter := 0
	n := 0
	packets := []packet{}
	for {
		p := packet{}
		if num > 0 && counter >= num {
			break
		}
		counter++
		if len(in[n:]) < 8 {
			break
		}
		p.ver = bitsToInt(in[n : n+3]...)
		n += 3
		p.typ = bitsToInt(in[n : n+3]...)
		n += 3
		if p.typ == 4 {
			// literal
			val := []int{}
			for in[n] != 0 {
				val = append(val, in[n+1:n+5]...)
				n += 5
			}
			val = append(val, in[n+1:n+5]...)
			p.val = bitsToInt(val...)
			n += 5
			packets = append(packets, p)
		} else {
			// operator
			lenTyp := in[n]
			n++
			if bitsToInt(lenTyp) == 0 {
				len := in[n : n+15]
				n += 15
				_, p.sub = parse(in[n:n+bitsToInt(len...)], -1)
				n += bitsToInt(len...)
			} else {
				num := in[n : n+11]
				n += 11
				nn := 0
				nn, p.sub = parse(in[n:], bitsToInt(num...))
				n += nn
			}
			packets = append(packets, p)
			packets = append(packets, p.sub...)
		}
	}
	return n, packets
}

func one(puzzle []byte) (string, error) {
	lines, err := common.SplitLines(puzzle)
	if err != nil {
		return "", err
	}
	for _, line := range lines[0] {
		tok := strings.Split(line, ",")
		bin := hexToBin(tok[0])
		_, packets := parse(bin, -1)
		tot := 0
		for _, p := range packets {
			tot += p.ver
		}
		fmt.Printf("%d==%s\t", tot, tok[1])
	}
	fmt.Printf("\n")
	return "", nil
}

func calc(packets []packet) int {
	out := 0
	for i := range packets {
		switch packets[i].typ {
		case 4: // literal
			out += packets[i].val
		case 0: // sum

		case 1: // product
		case 2: // minimum
		case 3: // maximum
		case 5: // gt
		case 6: // lt
		case 7: // eq

		}
	}
	return out
}

func two(puzzle []byte) (string, error) {
	lines, err := common.SplitLines(puzzle)
	if err != nil {
		return "", err
	}
	for _, line := range lines[0] {
		tok := strings.Split(line, ",")
		bin := hexToBin(tok[0])
		_, packets := parse(bin, -1)
		val := calc(packets)
		return fmt.Sprintf("%d", val), nil
	}
	return "", nil
}
