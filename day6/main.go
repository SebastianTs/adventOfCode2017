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
	for n := 0; ; n++ {
		cur := fmt.Sprint(balance(vs))
		if v, ok := seen[cur]; ok {
			return n + 1, n - v
		}
		seen[cur] = n
	}
}

func balance(vs []int) []int {
	maxIdx := getMaxIdx(vs)
	i := (maxIdx + 1) % len(vs)
	n := vs[maxIdx]
	vs[maxIdx] = 0
	for ; n > 0; n-- {
		vs[i]++
		i = (i + 1) % len(vs)
	}
	return vs
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
