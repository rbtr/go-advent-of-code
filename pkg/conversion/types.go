package conversion

import (
	"bufio"
	"bytes"
	"strings"
)

type Tokenizer func([]byte) [][]byte

type Parser[T any] func([]byte) (T, error)

var SplitLines Tokenizer = func(b []byte) [][]byte {
	sc := bufio.NewScanner(bytes.NewReader(b))
	out := [][]byte{}
	for sc.Scan() {
		line := make([]byte, len(sc.Bytes()))
		copy(line, sc.Bytes())
		out = append(out, line)
	}
	return out
}

var SplitCommas Tokenizer = func(b []byte) [][]byte {
	return StringsToBytes(strings.Split(string(b), ","))
}

var SplitCharacters Tokenizer = func(b []byte) [][]byte {
	sc := bufio.NewScanner(bytes.NewReader(b))
	sc.Split(bufio.ScanRunes)
	out := [][]byte{}
	for sc.Scan() {
		char := make([]byte, len(sc.Bytes()))
		copy(char, sc.Bytes())
		out = append(out, char)
	}
	return out
}

var SplitWords Tokenizer = func(b []byte) [][]byte {
	return StringsToBytes(strings.Split(string(b), " "))
}
