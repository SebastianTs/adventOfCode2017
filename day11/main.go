package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	steps := parseInput("./input")
	dist, max := distance(steps)
	fmt.Printf("The first distance is: %d\nThe max distance is: %d\n", dist, max)
}

func distance(steps []string) (dist, max int) {

	var x, y, z int = 0, 0, 0
	for _, step := range steps {
		switch step {
		case "n":
			y++
			z--
		case "ne":
			x++
			z--
		case "se":
			x++
			y--
		case "s":
			y--
			z++
		case "sw":
			x--
			z++
		case "nw":
			x--
			y++
		default:
		}
		dist = (abs(x) + abs(y) + abs(z)) / 2
		if max < dist {
			max = dist
		}
	}
	return
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func parseInput(file string) []string {
	out := make([]string, 0)

	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()

	s := bufio.NewScanner(fileHandle)

	for s.Scan() {
		line := s.Text()
		out = strings.Split(line, ",")
		break
	}
	return out
}
