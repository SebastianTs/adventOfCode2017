package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	set := parseInput("./input")
	vmax, vhigh := processInstructionSet(set)
	fmt.Printf("The largest value in any register is: \t%d\nThe highest value in any register was: \t%d\n", vmax, vhigh)
}
func processInstructionSet(in [][]string) (max, final int) {
	reg := make(map[string]int)
	cond := false
	for _, v := range in {
		if _, ok := reg[v[4]]; !ok {
			reg[v[4]] = 0
		}
		r := reg[v[4]]

		if r > max {
			max = r
		}

		x, err := strconv.Atoi(v[6])
		if err != nil {
			log.Fatal(err)
		}
		switch v[5] {
		case ">":
			cond = r > x
		case "<":
			cond = r < x
		case "<=":
			cond = r <= x
		case ">=":
			cond = r >= x
		case "==":
			cond = r == x
		case "!=":
			cond = r != x
		default:
			cond = false
		}
		if cond {
			if _, ok := reg[v[0]]; !ok {
				reg[v[0]] = 0
			}

			if reg[v[0]] > max {
				max = reg[v[0]]
			}
			y, err := strconv.Atoi(v[2])
			if err != nil {
				log.Fatal(err)
			}
			switch v[1] {
			case "inc":
				reg[v[0]] += y
			case "dec":
				reg[v[0]] -= y
			}
		}
	}
	for _, k := range reg {
		if k > final {
			final = k
		}
	}
	return final, max
}

func parseInput(file string) [][]string {
	out := make([][]string, 0)

	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()

	s := bufio.NewScanner(fileHandle)

	for s.Scan() {
		line := s.Text()
		t := bufio.NewScanner(strings.NewReader(line))
		t.Split(bufio.ScanWords)
		words := make([]string, 0)
		for t.Scan() {
			words = append(words, t.Text())
		}
		out = append(out, words)
	}
	return out
}
