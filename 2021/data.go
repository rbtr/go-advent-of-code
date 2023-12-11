package common

import (
	"bufio"
	"bytes"
	"strconv"
)

type Data []byte

func FromStrings(s ...string) Data {
	b := []byte{}
	for _, ss := range s {
		b = append(b, []byte(ss)...)
		b = append(b, '\n')
	}
	return b
}

// SplitLines returns a line-by-line string slice of the input []byte(s).
func (d Data) AsLines() []string {
	lines := []string{}
	s := bufio.NewScanner(bytes.NewReader(d))
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

// ParseMatrix parses a input bytes in to a matrix of ints.
func (d Data) AsIntMatrix() ([][]int, error) {
	sc := bufio.NewScanner(bytes.NewReader(d))
	sc.Split(bufio.ScanLines)
	lines := [][]int{}
	for sc.Scan() {
		sw := bufio.NewScanner(bytes.NewReader(sc.Bytes()))
		sw.Split(bufio.ScanRunes)
		line := []int{}
		for sw.Scan() {
			i, err := strconv.Atoi(sw.Text())
			if err != nil {
				return nil, err
			}
			line = append(line, i)
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func (d Data) AsStringMatrix() ([][]string, error) {
	sc := bufio.NewScanner(bytes.NewReader(d))
	sc.Split(bufio.ScanLines)
	lines := [][]string{}
	for sc.Scan() {
		sw := bufio.NewScanner(bytes.NewReader(sc.Bytes()))
		sw.Split(bufio.ScanRunes)
		line := []string{}
		for sw.Scan() {
			line = append(line, sw.Text())
		}
		lines = append(lines, line)
	}
	return lines, nil
}
