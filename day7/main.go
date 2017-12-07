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
	nodes := parseInput("./input_sorted")

	nodes = connect(nodes)
	fmt.Println(isRoot(nodes))
	fmt.Println(findInbalance(isRoot(nodes)))
	//fmt.Println(calcBalance(isRoot(nodes)))

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

func findInbalance(n node) (node, bool) {
	subtrees := make([]int, len(n.childs))
	help := make(map[int]int)
	for i := range n.childs {
		subtrees[i] = totalWeight(n.childs[i])
		help[subtrees[i]]++
	}
	for v, k := range help {
		if k == len(n.childs) {
			return node{}, false // subtree is balanced
		}
		if k == 1 {
			for i := range n.childs {
				if v == subtrees[i] { // TODO use map here
					return n.childs[i], true
				}
			}
		}

	}
	return node{}, false
}

func calcBalance(n node) int {
	pre := node{}
	found := false
	node := n
	for !found {
		pre = node
		node, found = findInbalance(n)
	}
	weight := 0
	for i, v := range pre.childs {
		weight = totalWeight(v)
		diff := totalWeight(pre.childs[(i+1)%len(pre.childs)]) - weight
		if diff != 0 {
			return diff
		}
	}

	return weight
}

func totalWeight(n node) int {
	weight := n.weight
	for i := range n.childs {
		weight += totalWeight(n.childs[i])
	}
	return weight
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

func connect(ns []node) []node {
	root := isRoot(ns)

	for i := range root.childs {
		if len(root.childs[i].childs) == 0 {
			for j := range ns {
				if ns[j].name == root.childs[i].name {
					root.childs[i] = ns[j]
				}
			}
		}
	}
	return ns
}

func (n node) print() {
	fmt.Println(n.name)
	for i := range n.childs {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		n.childs[i].print()
	}
}
