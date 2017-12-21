package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var grid = []string{".#.", "..#", "###"}

func main() {
	ins := parseInput("./input")
	_ = ins

}

func parseInput(file string) map[string]string {

	ins := make(map[string]string)
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()
	s := bufio.NewScanner(fileHandle)

	for s.Scan() {
		line := s.Text()
		t := strings.Split(line, " ")
		ins[t[0]] = t[2]
	}
	return ins
}
