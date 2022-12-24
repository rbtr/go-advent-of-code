package main

import (
	"fmt"
	"log"

	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/puzzle"
)

func main() {
	p, err := puzzle.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := p.Run(one, two); err != nil {
		log.Fatal(err)
	}
}

type operation struct {
	name, a, b, op string
}

func one(d puzzle.Raw) (string, error) {
	lines := conversion.SplitLines(d)

	numbers := map[string]int{}
	maths := map[string]operation{}
	for i := range lines {
		line := lines[i]
		name := line[:4]
		job := conversion.SplitWords(line[6:])
		if len(job) == 1 {
			num, _ := conversion.ParseInt(job[0])
			numbers[string(name)] = num
			continue
		}
		maths[string(name)] = operation{
			a:  string(job[0]),
			b:  string(job[2]),
			op: string(job[1]),
		}
	}

	for {
		if _, ok := numbers["root"]; ok {
			break
		}
		done := map[string]struct{}{}
		for name, v := range maths {
			var ok bool
			var a, b int
			if a, ok = numbers[v.a]; !ok {
				continue
			}
			if b, ok = numbers[v.b]; !ok {
				continue
			}
			res := 0
			switch v.op {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			numbers[name] = res
			done[name] = struct{}{}
		}
		for k := range done {
			delete(maths, k)
		}
	}
	return fmt.Sprintf("root: %d", numbers["root"]), nil
}

func two(d puzzle.Raw) (string, error) {
	lines := conversion.SplitLines(d)

	numbers := map[string]int{}
	maths := map[string]operation{}
	for i := range lines {
		line := lines[i]
		name := line[:4]
		if string(name) == "humn" {
			continue
		}
		job := conversion.SplitWords(line[6:])
		if len(job) == 1 {
			num, _ := conversion.ParseInt(job[0])
			numbers[string(name)] = num
			continue
		}
		op := string(job[1])
		if string(name) == "root" {
			op = "="
		}
		maths[string(name)] = operation{
			name: string(name),
			a:    string(job[0]),
			b:    string(job[2]),
			op:   op,
		}
	}
	equals := 0
	unsolved := ""
	apex := "root"
	for {
		if unsolved == "humn" {
			break
		}
		for {
			var ok bool
			if equals, ok = numbers[maths[apex].a]; ok {
				unsolved = maths[apex].b
				break
			}
			if equals, ok = numbers[maths[apex].b]; ok {
				unsolved = maths[apex].a
				break
			}
			done := map[string]struct{}{}
			for name, v := range maths {
				var ok bool
				var a, b int
				if a, ok = numbers[v.a]; !ok {
					continue
				}
				if b, ok = numbers[v.b]; !ok {
					continue
				}
				res := 0
				switch v.op {
				case "+":
					res = a + b
				case "-":
					res = a - b
				case "*":
					res = a * b
				case "/":
					res = a / b
				}
				numbers[name] = res
				done[name] = struct{}{}
			}
			for k := range done {
				delete(maths, k)
			}
		}
		for {
			var ok bool
			prev := unsolved
			math := maths[unsolved]
			fmt.Println(math)
			var a, b int
			if a, ok = numbers[math.a]; ok {
				if b, ok = numbers[math.b]; ok {
					break
				}
				unsolved = math.b
				switch math.op {
				case "+": // a + b = eq
					b = equals - a
				case "-": // a - b = eq
					b = a - equals
				case "*": // a * b = eq
					b = equals / a
				case "/": // a / b = eq
					b = a / equals
				}
				equals = b
				numbers[unsolved] = b
				delete(maths, prev)
				continue
			}
			if b, ok = numbers[math.b]; ok {
				unsolved = math.a
				switch math.op {
				case "+": // a + b = eq
					a = equals - b
				case "-": // a - b = eq
					a = equals + b
				case "*": // a * b = eq
					a = equals / b
				case "/": // a / b = eq
					a = equals * b
				}
				equals = a
				numbers[unsolved] = a
				delete(maths, prev)
				continue
			}
			apex = unsolved
			break
		}
		fmt.Println(numbers)
	}
	return fmt.Sprintf("humn: %d", equals), nil
}
