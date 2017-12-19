package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := parseInput("./input")
	result, steps := maze(input)
	fmt.Printf("The letters on the path were %s\nIt took %d steps.\n", result, steps)

}

func maze(in []string) (visited string, steps int) {
	x, y, dir, steps := 0, 0, 0, 0
	for j, c := range in[0] {
		if c == '|' {
			y = j
		}
	}

	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	size := len(dirs)
	for {
		cur := in[x][y]
		steps++
		switch cur {
		default:
			visited += string(cur)
			fallthrough
		case '-', '|':
			x += dirs[dir][0]
			y += dirs[dir][1]
		case '+':
			xd := (dir + 1) % size
			nx := dirs[xd]
			x2 := x + nx[0]
			y2 := y + nx[1]
			if in[x2][y2] != ' ' {
				dir, x, y = xd, x2, y2
				continue
			}

			yd := (size + dir - 1) % size
			ny := dirs[yd]
			x2 = x + ny[0]
			y2 = y + ny[1]
			if in[x2][y2] != ' ' {
				dir, x, y = yd, x2, y2
				continue
			}
		case ' ':
			return visited, steps - 1
		}

	}
}

func parseInput(file string) []string {
	val := make([]string, 0)
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()
	s := bufio.NewScanner(fileHandle)
	for s.Scan() {
		val = append(val, s.Text())
	}
	return val
}
