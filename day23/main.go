package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	ins := parseInput("./input")
	fmt.Printf("There were %d mul opertions.\n", process(ins))
	fmt.Printf("The value of h is %d. \n", primes())
}

func process(ins [][]string) int {

	regs := make(map[string]int)

	get := func(s string) int {
		if strings.IndexAny(s, "0123456789") != -1 {
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			return v
		}
		return regs[s]
	}
	var freq int
	for counter := 0; counter < len(ins); counter++ {
		in := ins[counter]
		switch in[0] {
		case "set":
			regs[in[1]] = get(in[2])
		case "sub":
			regs[in[1]] -= get(in[2])
		case "mul":
			regs[in[1]] *= get(in[2])
			freq++
		case "jnz":
			if get(in[1]) != 0 {
				counter += int(get(in[2]) - 1)
			}
		}
	}
	return freq
}

func primes() int64 {
	const start int64 = 79*100 + 1E5
	const end int64 = start + 17000
	var primes int64

	for i := int64(start); i != end+17; i += 17 {
		isPrime := 1

		for j := 2; j < int(math.Sqrt(float64(i)))+1; j++ {
			if i%int64(j) == 0 {
				isPrime = 0
				break
			}
		}
		if isPrime == 0 {
			primes++
		}
	}
	return primes
}

func parseInput(file string) [][]string {
	val := make([][]string, 0)
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()
	s := bufio.NewScanner(fileHandle)
	for s.Scan() {
		line := s.Text()
		token := strings.Split(line, " ")
		val = append(val, token)
	}

	return val
}
