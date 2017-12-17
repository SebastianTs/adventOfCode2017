package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	var input string
	if len(args) != 2 {
		input = "3"
		fmt.Printf("Use: %s puzzleinput\n", args[0])
		fmt.Printf("No Input was given will, use \"%s\" instead\n\n", input)
	} else {
		input = args[1]
	}
	steps, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The value next to 2017 is %d.\n", spinlock(steps))
	fmt.Printf("The value next to zero is %d.\n", spinlock5E7(steps))
}

func spinlock(steps int) int {
	buffer := make([]int, 1)
	cur := 0
	for i := 1; i < 2018; i++ {
		cur = ((cur + steps) % len(buffer)) + 1
		buffer = append(buffer[:cur], append([]int{i}, buffer[cur:]...)...)
	}
	return buffer[cur+1]
}

func spinlock5E7(steps int) int {
	cur := 0
	res := 0
	for i := 1; i < 5E7; i++ {
		cur = ((cur + steps) % i) + 1
		if cur == 1 {
			res = i
		}
	}
	return res
}
