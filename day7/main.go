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
	tree := parseInput("./input")
	root := root(tree)
	fmt.Printf("The root node is %s\n", root)
	size, wrongNode := balance(root, tree)
	fmt.Printf("The correct weight of %s is %d\n", wrongNode, size)
}

type graph struct {
	nodes map[string]int
	edges map[string][]string
}

func root(g graph) string {
	inc := make(map[string]bool)
	for _, list := range g.edges {
		for _, item := range list {
			inc[item] = true
		}
	}
	for node := range g.nodes {
		if !inc[node] {
			return node
		}
	}
	return ""
}

func weight(node string, g graph) int {
	w := g.nodes[node]
	for _, child := range g.edges[node] {
		w += weight(child, g)
	}
	return w
}

func findInbalance(node string, g graph) string {
	childs := make(map[int][]string)
	for _, child := range g.edges[node] {
		w := weight(child, g)
		childs[w] = append(childs[w], child)
	}
	for _, v := range childs {
		if len(v) == 1 {
			return v[0]
		}
	}
	return ""
}

func balance(node string, g graph) (int, string) {
	next := node
	visited := make([]string, 0)
	for next != "" {
		visited = append(visited, next)
		next = findInbalance(next, g)
	}
	last := len(visited) - 1
	for _, sibling := range g.edges[visited[last-1]] {
		if sibling != visited[last] {
			correct := weight(sibling, g)
			wrong := weight(visited[last], g)
			diff := correct - wrong
			return g.nodes[visited[last]] + diff, visited[last]
		}
	}
	return 0, ""
}

func drawDot(node string, g graph) string {
	s := "digraph day7 {\n"
	s += "rankdir=LR;\n"
	for k, v := range g.edges {
		for _, i := range v {
			s += fmt.Sprintf("%s [label=\"%s,%d\\n%d\"];", i, i, g.nodes[i], weight(i, g))
			s += fmt.Sprintf("%s -> %s;\n", k, i)
		}
	}
	s += "}"
	return s
}

func parseInput(file string) graph {
	nodes := make(map[string]int)
	edges := make(map[string][]string)

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
			edges[name] = token[3:]
		}
		if err != nil {
			log.Fatal(err)
		}
		nodes[name] = weight
	}
	return graph{nodes, edges}
}
