package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	values := parseInput("./input")
	fmt.Printf("First puzzle result:\t %d \n", leaveList(values))
	fmt.Printf("Second puzzle result:\t %d \n", leaveListStrange(values))
}

func parseInput(file string) []int {
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	s := bufio.NewScanner(fileHandle)

	values := make([]int, 0)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, i)
	}
	return values
}

func leaveList(list []int) int {
	count, cur := 0, 0
	for {
		next := list[cur] + cur
		count++
		if next >= len(list) || next < 0 {
			return count
		}
		list[cur]++
		cur = next
	}
}

func leaveListStrange(list []int) int {
	count, cur, next := 0, 0, 0
	for {
		offset := list[cur]
		next += offset
		if offset >= 3 {
			list[cur]--
		} else {
			list[cur]++
		}
		cur = next
		count++
		if next >= len(list) || next < 0 {
			return count
		}
	}

}
