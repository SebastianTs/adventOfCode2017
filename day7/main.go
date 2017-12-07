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
	nodes := parseInput("./input")
	fmt.Printf("The rootnode is: %s\n", isRoot(nodes).name)
}

type node struct {
	name   string
	weight int
	childs []node
}

func isRoot(nodes []node) node {
	incoming := make(map[string]bool)
	for _, v := range nodes {
		for _, n := range v.childs {
			incoming[n.name] = true
		}
	}
	for _, v := range nodes {
		if _, ok := incoming[v.name]; !ok {
			return v
		}
	}
	return node{}
}

func parseInput(file string) []node {
	out := make([]node, 0)
	adj := make(map[string][]string)

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
			token = append(token, strings.Trim(t.Text(), "(),"))
		}
		name := token[0]
		weight, err := strconv.Atoi(token[1])

		if len(token) > 3 {
			adj[name] = token[3:]
		}
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, node{name, weight, nil})
	}

	for i := range out {
		for node, adjList := range adj {
			if node == out[i].name {
				for j := range out {
					for _, n := range adjList {
						if n == out[j].name {
							out[i].childs = append(out[i].childs, out[j])
						}
					}
				}
			}
		}
	}
	return out
}
