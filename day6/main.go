package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	vs := parseInput("./input")
	index, steps := hasCycle(vs)
	fmt.Printf("Number of cycles:\t %d \nLoop size:\t\t %d\n", index, steps)
}

func hasCycle(vs []int) (count, cycles int) {
	seen := make(map[string]int)
	n := 0
	for {
		cur := fmt.Sprint(balance(vs))
		if v, ok := seen[cur]; ok {
			return n + 1, n - v
		}
		seen[cur] = n
		n++
	}
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func getMaxIdx(vs []int) (max int) {
	maxV := vs[0]
	for i, v := range vs[1:] {
		if v > maxV {
			maxV = v
			max = i + 1
		}
	}
	return
}

func balance(vs []int) []int {
	maxIdx := getMaxIdx(vs)
	i := (maxIdx + 1) % len(vs)
	n := vs[maxIdx]
	vs[maxIdx] = 0
	for n > 0 {
		vs[i]++
		n--
		i = (i + 1) % len(vs)
	}
	return vs

}

func parseInput(file string) []int {
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	s := bufio.NewScanner(fileHandle)
	s.Split(bufio.ScanWords)
	values := make([]int, 0)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, i)
	}
	return values
}
