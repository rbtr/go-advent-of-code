package main

import (
	"bufio"
	"bytes"
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
	if err := one(puzzle); err != nil {
		log.Fatal(err)
	}
	if err := two(puzzle); err != nil {
		log.Fatal(err)
	}
}

func one(puzzle *common.Puzzle) error {
	sc := bufio.NewScanner(bytes.NewReader(puzzle.Input))
	sc.Split(bufio.ScanLines)
	sc.Scan()
	splits := strings.Split(sc.Text(), ",")
	arr := make([]int, len(splits))
	var err error
	for i := range splits {
		if arr[i], err = strconv.Atoi(splits[i]); err != nil {
			return err
		}
	}
	days := 80
	for i := 0; i < days; i++ {
		new := 0
		for j := range arr {
			arr[j] -= 1
			if arr[j] < 0 {
				arr[j] = 6
				new++
			}
		}
		for j := 0; j < new; j++ {
			arr = append(arr, 8)
		}
	}
	fmt.Printf("%d\n", len(arr))
	return nil
}

func two(puzzle *common.Puzzle) error {
	sc := bufio.NewScanner(bytes.NewReader(puzzle.Input))
	sc.Split(bufio.ScanLines)
	sc.Scan()
	splits := strings.Split(sc.Text(), ",")
	buckets := make([]int, 9)
	for i := range splits {
		ii, err := strconv.Atoi(splits[i])
		if err != nil {
			return err
		}
		buckets[ii]++
	}
	fmt.Println(buckets)

	days := 256
	for i := 0; i < days; i++ {
		new := buckets[0]
		for j := 0; j < 8; j++ {
			buckets[j] = buckets[j+1]
		}
		buckets[6] += new
		buckets[8] = new
	}
	total := 0
	for _, n := range buckets {
		total += n
	}
	fmt.Printf("%v %d\n", buckets, total)
	return nil
}
