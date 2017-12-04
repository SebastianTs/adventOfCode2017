package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	passphrases := parseInput("./input")

	fmt.Printf("First puzzle result: %d \n", countValid(isValid, passphrases))
	fmt.Printf("Second puzzle result: %d \n", countValid(isValidWithoutAnagram, passphrases))

}

func countValid(fn func([]string) bool, input [][]string) (sum int) {
	for _, words := range input {
		if fn(words) {
			sum++
		}
	}
	return
}

func isValid(s []string) bool {
	count := make(map[string]bool)
	for _, e := range s {
		if count[e] {
			return false
		}
		count[e] = true
	}
	return true
}

func isValidWithoutAnagram(s []string) bool {
	count := make(map[string]bool)
	for _, e := range s {
		cur := sortString(e)
		if count[cur] {
			return false
		}
		count[cur] = true
	}
	return true
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func parseInput(file string) [][]string {
	out := make([][]string, 0)

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
		words := make([]string, 0)
		for t.Scan() {
			words = append(words, t.Text())
		}
		out = append(out, words)
	}
	return out
}
