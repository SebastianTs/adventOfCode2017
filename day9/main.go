package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	stream := parseInput("./input")
	score, garbageCount := countGarbage(stream)
	fmt.Printf("The score is %d.\nCounted %d characters as garbage.\n", score, garbageCount)

}

func countGarbage(stream string) (int, int) {
	var garbage bool
	var c byte
	var score, depth, garbageCount int = 0, 0, 0

	for i := 0; i < len(stream); i++ {
		c = stream[i]
		if c == '!' {
			i++
		} else if garbage && c != '>' {
			garbageCount++
		} else if c == '<' {
			garbage = true
		} else if c == '>' {
			garbage = false
		} else if c == '{' {
			depth++
			score += depth
		} else if c == '}' {
			depth--

		}
	}
	return score, garbageCount
}

func parseInput(file string) string {
	out := ""

	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()

	s := bufio.NewScanner(fileHandle)

	for s.Scan() {
		line := s.Text()
		out += line
	}
	return out
}
