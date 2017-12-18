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
	ins := parseInput("./input")
	fmt.Println(process(ins))
}

func process(ins [][]string) int64 {

	regs := make(map[string]int64)

	get := func(s string) int64 {
		if strings.IndexAny(s, "0123456789") != -1 {
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			return int64(v)
		}
		return regs[s]
	}
	found := false
	var freq int64
	for counter := 0; !found && counter < len(ins); counter++ {
		in := ins[counter]
		switch in[0] {
		case "snd":
			freq = regs[in[1]]
		case "set":
			regs[in[1]] = get(in[2])
		case "add":
			regs[in[1]] += get(in[2])
		case "mul":
			regs[in[1]] *= get(in[2])
		case "mod":
			regs[in[1]] %= get(in[2])
		case "rcv":
			if get(in[1]) != 0 {
				regs[in[1]] = freq
				found = true
			}
		case "jgz":
			if regs[in[1]] > 0 {
				counter += int(get(in[2]) - 1)
			}
		}
	}
	return freq
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
