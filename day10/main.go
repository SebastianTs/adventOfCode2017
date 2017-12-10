package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const listSize = 256

func main() {
	seed := parseInput("./input")
	list := knot(seed)
	fmt.Printf("The solution to puzzle part1 is: %d \n", (list[0] * list[1]))
	s := parseInputByte("./input")
	fmt.Printf("The solution to puzzle part2 is: %x \n", knotByte(s))

}

func knot(ls []int) []int {
	list := make([]int, listSize)
	for i := range list {
		list[i] = i
	}
	cur := 0
	skip := 0
	for _, l := range ls {
		list = revSubList(cur, l, list)
		cur = (cur + l + skip) % len(list)
		skip++
	}
	return list
}

func revSubList(s, l int, list []int) []int {
	for i := 0; i < l/2; i++ {
		left := (i + s) % len(list)
		right := (l - 1 - i + s) % len(list)
		list[left], list[right] = list[right], list[left]

	}
	return list
}

func parseInput(file string) []int {
	out := (make([]int, 0))
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()

	s := bufio.NewScanner(fileHandle)
	for s.Scan() {
		line := s.Text()
		vals := strings.Split(line, ",")
		for _, v := range vals {
			value, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			out = append(out, value)
		}
	}
	return out
}

func parseInputByte(file string) []byte {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	for _, v := range []byte{17, 31, 73, 47, 23} {
		b = append(b, v)
	}
	return b
}

func knotByte(ls []byte) []byte {
	list := make([]byte, listSize)
	for i := range list {
		list[i] = byte(i)
	}
	cur := 0
	skip := 0
	for i := 0; i < 64; i++ {
		for _, l := range ls {
			list = revSubListByte(cur, int(l), list)
			cur = (cur + int(l) + skip) % len(list)
			skip++
		}
	}
	list = sToDense(list)
	return list
}

func sToDense(sparse []byte) (dense []byte) {
	for i := 0; i < 16; i++ {
		dense = append(dense, sToDenseShort(sparse[i*16:]))
	}
	return dense
}

func sToDenseShort(sparse []byte) (dense byte) {
	var res byte
	for j := 0; j < 16; j++ {
		res ^= sparse[j]
	}
	return res
}

func revSubListByte(s, l int, list []byte) []byte {
	for i := 0; i < l/2; i++ {
		left := (i + s) % len(list)
		right := (l - 1 - i + s) % len(list)
		list[left], list[right] = list[right], list[left]

	}
	return list
}
