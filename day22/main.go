package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	node, center := parseInput("./input")
	fmt.Println(sporifica(node, center, 1E4))
	fmt.Println(sporificaEnhanced(node, center, 1E7))
}

func sporifica(node map[[2]int]bool, center [2]int, bursts int) (count int) {
	//			      Up      Right   Down     Left
	dirs := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	dir := 0
	cur := center
	for i := 0; i < bursts; i++ {
		//print(node, center, cur)
		if node[cur] {
			dir = (dir + 1) % 4
		} else {
			dir = (dir + 3) % 4
			count++
		}
		node[cur] = !node[cur]
		cur[0] += dirs[dir][0]
		cur[1] += dirs[dir][1]
	}
	return count
}

func sporificaEnhanced(node map[[2]int]bool, center [2]int, bursts int) (count int) {
	dirs := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	dir := 0
	cur := center
	enode := make(map[[2]int]rune)
	for xy := range node {
		enode[xy] = '#'
	}
	for i := 0; i < bursts; i++ {
		//printEnhanced(enode, center, cur)
		switch enode[cur] {
		default: // Clean
			dir = (dir + 3) % 4
			enode[cur] = 'W'
		case '#':
			dir = (dir + 1) % 4
			enode[cur] = 'F'
		case 'W':
			count++
			enode[cur] = '#'
		case 'F':
			dir = (dir + 2) % 4
			enode[cur] = '.'
		}
		cur[0] += dirs[dir][0]
		cur[1] += dirs[dir][1]
	}
	return count
}

func print(node map[[2]int]bool, center [2]int, cur [2]int) {
	for m := -4; m < 6; m++ {
		for n := -4; n < 6; n++ {
			pos := [2]int{n, m}
			s := ""
			if node[pos] {
				s = "#"
			} else {
				s = "."
			}
			if pos == cur {
				fmt.Printf("[%s]", s)
			} else {

				fmt.Printf(" %s ", s)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("")
}

func printEnhanced(node map[[2]int]rune, center [2]int, cur [2]int) {
	for m := -4; m < 6; m++ {
		for n := -4; n < 6; n++ {
			pos := [2]int{n, m}
			s := ""
			if v, ok := node[pos]; ok {
				s = string(v)
			} else {
				s = "."
			}
			if pos == cur {
				fmt.Printf("[%s]", s)
			} else {

				fmt.Printf(" %s ", s)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("")
}

func parseInput(file string) (node map[[2]int]bool, center [2]int) {
	node = make(map[[2]int]bool)
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()
	s := bufio.NewScanner(fileHandle)
	i, j := 0, 0
	for s.Scan() {
		line := s.Text()
		for k, c := range line {
			switch c {
			case '#':
				node[[2]int{k, i}] = true
			case '.':
				//node[[2]int{j, i}] = false
			}
			j = k
		}
		i++
	}
	return node, [2]int{j / 2, i / 2}
}
