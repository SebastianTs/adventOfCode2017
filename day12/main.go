package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	gr "github.com/yourbasic/graph"
)

func main() {
	graph := parseInput("./input")
	zeroSize, size := lazySolve(graph)
	fmt.Printf("The group that program 0 can talk to has %d programs.\n", zeroSize)
	fmt.Printf("There are %d groups of programs\n", size)
}

type graph struct {
	nodes []int
	edges [][2]int
}

func lazySolve(g graph) (maxSize, size int) {
	graph := gr.New(len(g.nodes))
	for _, e := range g.edges {
		graph.AddBoth(e[0], e[1])
	}

	comp := gr.StrongComponents(graph)
	max := 0
	for _, list := range comp {
		for _, item := range list {
			if item == 0 {
				max = len(list)
			}
		}
	}
	return max, len(comp)
}

func parseInput(file string) graph {
	nodes := make([]int, 0)
	edges := make([][2]int, 0)

	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()

	s := bufio.NewScanner(fileHandle)

	for s.Scan() {
		line := s.Text()
		t := bufio.NewScanner(strings.NewReader(line))
		t.Split(bufio.ScanWords)
		token := make([]string, 0)
		for t.Scan() {
			token = append(token, strings.Trim(t.Text(), ","))
		}

		n, err := strconv.Atoi(token[0])
		if err != nil {
			log.Fatal(err)
		}
		nodes = append(nodes, n)
		for _, t := range token[2:] {
			edge, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal(err)
			}
			edges = append(edges, [2]int{n, edge})
		}

	}
	return graph{nodes, edges}
}
