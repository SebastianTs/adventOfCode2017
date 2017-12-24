package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	cs                 []component
	mStrength, mLength int
)

type component struct {
	port0 int
	port1 int
	inUse bool
}

func main() {
	cs = parseInput("./input")
	fmt.Printf("The strength of the strongest bridge one can make is %d.\n", strengthOfBridge(buildBridge))
	fmt.Printf("The strength of the longest bridge one can make is %d.\n", strengthOfBridge(buildLongestBridge))
}

func strengthOfBridge(fn func(p, s, l int)) int {
	mStrength = -1
	mLength = -1
	fn(0, 0, 0)
	return mStrength
}

func buildBridge(port, strength, l int) {
	if mStrength < strength {
		mStrength = strength
		mLength = l
	}
	for i, c := range cs {
		if !c.inUse {
			cur := strength + c.port0 + c.port1
			switch port {
			case c.port0:
				cs[i].inUse = true
				buildBridge(c.port1, cur, l+1)
				cs[i].inUse = false

			case c.port1:
				cs[i].inUse = true
				buildBridge(c.port0, cur, l+1)
				cs[i].inUse = false
			}
		}
	}
}

func buildLongestBridge(port, strength, l int) {
	if mStrength < strength && mLength <= l || mLength < l {
		mStrength = strength
		mLength = l
	}
	for i, c := range cs {
		if !c.inUse {
			cur := strength + c.port0 + c.port1
			switch port {
			case c.port0:
				cs[i].inUse = true
				buildLongestBridge(c.port1, cur, l+1)
				cs[i].inUse = false
			case c.port1:
				cs[i].inUse = true
				buildLongestBridge(c.port0, cur, l+1)
				cs[i].inUse = false
			}
		}
	}
}

func parseInput(file string) []component {
	out := make([]component, 0)
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(2)
	}
	defer fileHandle.Close()
	s := bufio.NewScanner(fileHandle)
	for s.Scan() {
		tokens := strings.Split(s.Text(), "/")
		port0, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatal(err)
		}
		port1, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal(err)
		}
		out = append(out,
			component{
				port0: port0,
				port1: port1,
				inUse: false})
	}
	return out
}
