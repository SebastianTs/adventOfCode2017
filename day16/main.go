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
	moves := parseInput("./input")

	progs := [16]rune{}
	for i := 0; i < len(progs); i++ {
		progs[i] = rune('a' + i)
	}
	dance(&progs, moves)
	fmt.Printf("The order of dancers after the first dance is\t %s.\n", dancers(&progs))
	billionDances(&progs, moves)
	fmt.Printf("The order of dancers after a billion dances is\t %s.\n", dancers(&progs))

}

func billionDances(progs *[16]rune, moves []string) {
	c, perms := cycle(progs, moves)
	steps := 1E10 % c
	for k, v := range perms {
		if v == steps {
			*progs = k
			break
		}
	}
}

func cycle(progs *[16]rune, moves []string) (int, map[[16]rune]int) {
	seen := make(map[[16]rune]int)
	i := 1
	for {
		seen[*progs] = i
		dance(progs, moves)
		if _, ok := seen[*progs]; ok {
			break
		}
		i++
	}
	return i, seen
}

func dance(progs *[16]rune, moves []string) {
	for _, move := range moves {
		switch move[0] {
		case 's':
			steps := parseS(move[1:])
			spin(progs, steps)
		case 'x':
			a, b := parseX(move[1:])
			exchange(progs, a, b)
		case 'p':
			a, b := parseP(move[1:])
			partner(progs, a, b)
		}
	}
}

func spin(progs *[16]rune, steps int) {
	left := make([]rune, 16)
	right := make([]rune, 16)
	left = progs[16-steps:]
	right = progs[:16-steps]
	res := append(left, right...)
	for i := range progs {
		progs[i] = res[i]
	}
}

func exchange(progs *[16]rune, a, b int) {
	progs[a], progs[b] = progs[b], progs[a]
}

func partner(progs *[16]rune, a, b rune) {
	var posA, posB int
	for i, c := range progs {
		if c == a {
			posA = i
		}
		if c == b {
			posB = i
		}
	}
	exchange(progs, posA, posB)
}

func parseS(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func parseX(input string) (int, int) {
	foo := strings.Split(input, "/")
	a, err := strconv.Atoi(foo[0])
	b, err := strconv.Atoi(foo[1])
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}

func parseP(input string) (rune, rune) {
	foo := strings.Split(input, "/")
	return rune(foo[0][0]), rune(foo[1][0])
}

func parseInput(file string) []string {
	var token []string
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()
	s := bufio.NewScanner(fileHandle)
	for s.Scan() {
		line := s.Text()
		token = strings.Split(line, ",")
		break
	}
	return token
}

func dancers(progs *[16]rune) string {
	s := ""
	for i := range progs {
		s += string(progs[i])
	}
	return s
}
