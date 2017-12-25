package main

import "fmt"

func main() {

	state := 'A'
	slot := 0
	tape := make(map[int]bool)

	for steps := 12656374; steps > 0; steps-- {
		switch state {
		case 'A':
			if !tape[slot] {
				tape[slot] = true
				slot++
				state = 'B'
			} else {
				tape[slot] = false
				slot--
				state = 'C'
			}
		case 'B':
			if !tape[slot] {
				tape[slot] = true
				slot--
				state = 'A'
			} else {
				tape[slot] = true
				slot--
				state = 'D'
			}
		case 'C':
			if !tape[slot] {
				tape[slot] = true
				slot++
				state = 'D'
			} else {
				tape[slot] = false
				slot++
				state = 'C'
			}
		case 'D':
			if !tape[slot] {
				tape[slot] = false
				slot--
				state = 'B'
			} else {
				tape[slot] = false
				slot++
				state = 'E'
			}
		case 'E':
			if !tape[slot] {
				tape[slot] = true
				slot++
				state = 'C'
			} else {
				tape[slot] = true
				slot--
				state = 'F'
			}
		case 'F':
			if !tape[slot] {
				tape[slot] = true
				slot--
				state = 'E'
			} else {
				tape[slot] = true
				slot++
				state = 'A'
			}
		}
	}
	counter := 0
	for _, v := range tape {
		if v {
			counter++
		}
	}
	fmt.Printf("The checksum is: %d\n", counter)
}
