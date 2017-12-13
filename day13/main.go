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
	m := parseInput("./input")
	fmt.Printf("The severity of the trip is %d.\n", firewallState(m))
	fmt.Printf("The delay needs to be %d.\n", firewallStateDelay(m))
}

func firewallState(m map[int]int) int {
	res := 0
	for steps, size := range m {
		if steps%(2*size-2) == 0 {
			res += steps * size
		}
	}
	return res
}

func firewallStateDelay(m map[int]int) int {

	delay := 0
	for {
		caught := false
		for steps, size := range m {
			if (steps+delay)%(2*size-2) == 0 {
				caught = true
				break
			}
		}
		if !caught {
			return delay
		}
		delay++
	}
}

func parseInput(file string) map[int]int {
	val := make(map[int]int)
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
		t := len(token[0])
		idx, err := strconv.Atoi(token[0][:t-1])
		if err != nil {
			log.Fatal(err)
		}
		size, err := strconv.Atoi(token[1])
		if err != nil {
			log.Fatal(err)
		}
		val[idx] = size
	}

	return val
}
