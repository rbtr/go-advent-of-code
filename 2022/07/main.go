package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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

type node struct {
	name     string
	parent   *node
	children []*node
	size     int
}

func (n *node) padString(depth int) string {
	s := ""
	for i := 0; i < depth; i++ {
		s += "-"
	}
	s += fmt.Sprintf("%s DIR=%t Size=%d\n", n.name, n.isDir(), n.size)
	for i := range n.children {
		s += n.children[i].padString(depth + 1)
	}
	return s
}

func (n *node) String() string {
	s := fmt.Sprintf("\t%s\n\t", n.name)
	for i := range n.children {
		s += n.children[i].String()
	}
	return n.padString(0)
}

func (n *node) isDir() bool {
	return len(n.children) > 0
}

func (n *node) getSize() int {
	if n.size == 0 {
		for i := range n.children {
			n.size += n.children[i].getSize()
		}
	}
	return n.size
}

func parseTree(d []byte) *node {
	lines := conversion.SplitLines(d)
	root := &node{name: "/"}
	current := &node{
		children: []*node{root},
	}
	for i := range lines {
		tokens := strings.Split(string(lines[i]), " ")
		switch tokens[0] {
		case "$":
			switch tokens[1] {
			case "cd":
				switch tokens[2] {
				case "..":
					current = current.parent
				default:
					for j := range current.children {
						if current.children[j].name == tokens[2] {
							current = current.children[j]
							break
						}
					}
				}
			}
		case "dir":
			dir := &node{
				name:   tokens[1],
				parent: current,
			}
			current.children = append(current.children, dir)
		default:
			size, _ := strconv.Atoi(tokens[0])
			file := &node{
				name:   tokens[1],
				parent: current,
				size:   size,
			}
			current.children = append(current.children, file)
		}
	}
	return root
}

func sizeDirs(root *node) int {
	sum := 0
	if root.getSize() < 100000 {
		sum += root.getSize()
	}
	for i := range root.children {
		if root.children[i].isDir() {
			sum += sizeDirs(root.children[i])
		}
	}
	return sum
}

func one(d puzzle.Raw) (string, error) {
	tree := parseTree(d)
	sum := sizeDirs(tree)
	return fmt.Sprintf("%d", sum), nil
}

func findDir(root *node, best *node, target int) *node {
	if root.getSize() > target {
		if root.getSize() < best.getSize() {
			best = root
		}
	}
	for i := range root.children {
		if root.children[i].isDir() {
			best = findDir(root.children[i], best, target)
		}
	}
	return best
}

func two(d puzzle.Raw) (string, error) {
	total := 70000000
	need := 30000000

	tree := parseTree(d)

	used := tree.getSize()
	free := total - used

	tofree := need - free
	best := findDir(tree, tree, tofree)
	return fmt.Sprintf("%d", best.getSize()), nil
}
