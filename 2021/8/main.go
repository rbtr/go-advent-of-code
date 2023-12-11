// nolint
package main

import (
	"bufio"
	"bytes"
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
	// if err := one(puzzle); err != nil {
	// 	log.Fatal(err)
	// }
	if err := two(puzzle); err != nil {
		log.Fatal(err)
	}
}

// to render this takes this many segments
const (
	oneDigit   = 2
	sevenDigit = 3
	fourDigit  = 4
	eightDigit = 7
)

func one(puzzle *common.Puzzle) error {
	sc := bufio.NewScanner(bytes.NewReader(puzzle.Sample))
	sc.Split(bufio.ScanLines)
	lines := [][]string{}
	for sc.Scan() {
		line := strings.Split(sc.Text(), "|")
		lines = append(lines, line)
	}

	total := 0
	for i := range lines {
		outputs := lines[i][1]
		sw := bufio.NewScanner(bytes.NewReader([]byte(outputs)))
		sw.Split(bufio.ScanWords)
		for sw.Scan() {
			w := sw.Text()
			switch len(w) {
			case oneDigit, sevenDigit, fourDigit, eightDigit:
				total++
			}
		}
		strings.Split(outputs, " ")
	}
	fmt.Println(total)
	return nil
}

// var (
// 	zeroBits  = "1110111" // six
// 	oneBits   = "0010010" // known
// 	twoBits   = "1011101" // five
// 	threeBits = "1011011" // five
// 	fourBits  = "0111010" // known
// 	fiveBits  = "1101011" // five
// 	sixBits   = "1101111" // six
// 	sevenBits = "1010010" // known
// 	eightBits = "1111111" // known
// 	nineBits  = "1111011" // six
// )

// abcdefg
// 1101000

// 1010010
// 0010010

// 1101000 dab 7
// 1100000  ab 1
// 0001000 d   => a

// dab - ab = d -> 1 (a)

// 8 - 4 - a

// 8 - 4 - (a) = 8 - 0 = (d)
// 8 - 4 - (a)       (d)

// 8 - 6 = (c)
// 8 - 9 = (e)

// a = 7 | 1
// eg = 8 | 4 | a

// (2) = (5) | eg == 2
func xor(a, b []bool) []bool {
	out := make([]bool, len(a))
	for i := range a {
		out[i] = a[i] != b[i]
	}
	return out
}

// func varxor(args ...[]bool) []bool {
// 	out := args[0]
// 	for i := 1; i < len(args); i++ {
// 		for j := range out {
// 			out[j] = out[j] != args[i][j]
// 		}
// 	}
// 	return out
// }

func sumBits(b []bool) int {
	sum := 0
	for i := range b {
		if b[i] {
			sum++
		}
	}
	return sum
}

func eqBits(a, b []bool) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func toBits(s string) []bool {
	out := make([]bool, 7)
	for i := range s {
		switch s[i] {
		case 'a':
			out[0] = true
		case 'b':
			out[1] = true
		case 'c':
			out[2] = true
		case 'd':
			out[3] = true
		case 'e':
			out[4] = true
		case 'f':
			out[5] = true
		case 'g':
			out[6] = true
		}
	}
	return out
}

func two(puzzle *common.Puzzle) error {
	sc := bufio.NewScanner(bytes.NewReader(puzzle.Input))
	sc.Split(bufio.ScanLines)
	lines := [][]string{}
	for sc.Scan() {
		line := strings.Split(sc.Text(), "|")
		lines = append(lines, line)
	}
	total := 0
	for i := range lines {
		sw := bufio.NewScanner(bytes.NewReader([]byte(lines[i][0])))
		sw.Split(bufio.ScanWords)

		bits := make([][]bool, 10)
		unknowns := [][]bool{}
		for sw.Scan() {
			w := sw.Text()
			switch len(w) {
			case oneDigit:
				bits[1] = toBits(w)
			case fourDigit:
				bits[4] = toBits(w)
			case sevenDigit:
				bits[7] = toBits(w)
			case eightDigit:
				bits[8] = toBits(w)
			default:
				unknowns = append(unknowns, toBits(w))
			}
		}

		eg := xor(xor(bits[8], bits[7]), xor(bits[4], bits[1]))
		for i := range unknowns {
			if sumBits(xor(eg, unknowns[i])) == 3 {
				bits[2] = unknowns[i]
			}
		}

		for i := range unknowns {
			sum := sumBits(xor(bits[2], unknowns[i]))
			if sum == 2 {
				bits[3] = unknowns[i]
			}
			if sum == 4 {
				bits[5] = unknowns[i]
			}
		}

		for i := range unknowns {
			sum := sumBits(xor(bits[7], unknowns[i]))
			if sum == 5 {
				bits[6] = unknowns[i]
			}
		}

		for i := range unknowns {
			sum := sumBits(xor(bits[3], unknowns[i]))
			if sum == 1 {
				bits[9] = unknowns[i]
			}
		}

		for i := range unknowns {
			sum := sumBits(xor(xor(bits[7], bits[4]), unknowns[i]))
			if sum == 5 {
				bits[0] = unknowns[i]
			}
		}

		sww := bufio.NewScanner(bytes.NewReader([]byte(lines[i][1])))
		sww.Split(bufio.ScanWords)
		digits := []string{}
		for sww.Scan() {
			bit := toBits(sww.Text())
			for i := range bits {
				if eqBits(bit, bits[i]) {
					digits = append(digits, strconv.Itoa(i))
				}
			}
		}

		res := strings.Join(digits, "")
		this, err := strconv.Atoi(res)
		if err != nil {
			return err
		}
		total += this
	}
	fmt.Println(total)
	return nil
	// decode := map[string]int
	// return nil
}

func prettyPrint(b [][]bool) {
	for i := range b {
		fmt.Printf("%d: %v\n", i, b[i])
	}
}

func prettyPrintLookups(m map[string]int) {
	for k, v := range m {
		fmt.Printf("%d: %s\n", v, k)
	}
}
