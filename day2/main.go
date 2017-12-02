package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	table := parseInput("./input")

	result := checksum(rowChecksum, table)
	fmt.Printf("First puzzle result: %d \n", result)

	result = checksum(rowChecksumEven, table)
	fmt.Printf("Second puzzle result: %d \n", result)
}

func checksum(fn func([]int) int, table [][]int) (sum int) {
	for _, row := range table {
		sum += fn(row)
	}
	return
}

func rowChecksum(row []int) int {
	min := 1<<63 - 1 // MaxInt64
	max := -1 << 63  // MinInt64

	for _, el := range row {
		if el < min {
			min = el
		}
		if el > max {
			max = el
		}
	}
	return max - min
}

func rowChecksumEven(row []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(row)))

	for i := 0; i < len(row)-1; i++ {
		for j := i + 1; j < len(row); j++ {
			if row[i]%row[j] == 0 {
				return row[i] / row[j]
			}
		}
	}
	return 0
}

func parseInput(file string) [][]int {
	out := make([][]int, 0)

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
		values := make([]int, 0)
		for t.Scan() {
			i, err := strconv.Atoi(t.Text())
			if err != nil {
				log.Fatal(err)
			}
			values = append(values, i)
		}
		out = append(out, values)
	}
	return out
}
