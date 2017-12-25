package main

import "fmt"

func main() {

	state := 0
	slot := 0
	tape := make(map[int]int)

	for steps := 12656374; steps > 0; steps-- {
		switch state {
		case 0:
			if tape[slot] == 0 {
				tape[slot] = 1
				slot++
				state = 1
			} else if tape[slot] == 1 {
				tape[slot] = 0
				slot--
				state = 2
			}
		case 1:
			if tape[slot] == 0 {
				tape[slot] = 1
				slot--
				state = 0
			} else if tape[slot] == 1 {
				tape[slot] = 1
				slot--
				state = 3
			}
		case 2:
			if tape[slot] == 0 {
				tape[slot] = 1
				slot++
				state = 3
			} else if tape[slot] == 1 {
				tape[slot] = 0
				slot++
				state = 2
			}
		case 3:
			if tape[slot] == 0 {
				tape[slot] = 0
				slot--
				state = 1
			} else if tape[slot] == 1 {
				tape[slot] = 0
				slot++
				state = 4
			}
		case 4:
			if tape[slot] == 0 {
				tape[slot] = 1
				slot++
				state = 2
			} else if tape[slot] == 1 {
				tape[slot] = 1
				slot--
				state = 5
			}
		case 5:
			if tape[slot] == 0 {
				tape[slot] = 1
				slot--
				state = 4
			} else if tape[slot] == 1 {
				tape[slot] = 1
				slot++
				state = 0
			}
		}
	}
	counter := 0
	for _, v := range tape {
		if v == 1 {
			counter++
		}
	}
	fmt.Printf("The checksum is: %d\n", counter)
}
