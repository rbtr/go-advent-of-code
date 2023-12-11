package common

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

// Ingest reads a file in to memory.
func Ingest(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// SplitLines returns a line-by-line string slice of the input []byte(s).
func SplitLines(b ...[]byte) ([][]string, error) {
	lines := make([][]string, len(b))
	for i := range b {
		s := bufio.NewScanner(bytes.NewReader(b[i]))
		for s.Scan() {
			lines[i] = append(lines[i], s.Text())
		}
	}
	return lines, nil
}

// ParseMatrix parses a input bytes in to a matrix of ints.
func ParseMatrix(b []byte) ([][]int, error) {
	sc := bufio.NewScanner(bytes.NewReader(b))
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
