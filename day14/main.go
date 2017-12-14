package main

import (
	"fmt"
	"math/bits"
	"os"
)

func main() {
	args := os.Args
	var input string
	if len(args) != 2 {
		input = "AoC2017"
		fmt.Printf("Use: %s puzzleinput\n", args[0])
		fmt.Printf("No Input was given will, use \"%s\" instead\n\n", input)
	} else {
		input = args[1]
	}
	fmt.Printf("There are %d squares used.\n", defragWithKnot(input))
	grid := buildGrid(input)
	fmt.Printf("Threr are %d regions used.\n", countRegions(grid))

}

func defragWithKnot(s string) (count int) {
	for i := 0; i < 128; i++ {
		cur := fmt.Sprintf("%s-%d", s, i)
		row := knotByte([]byte(cur))
		count += countOnes(row)
	}
	return
}

func countOnes(in []byte) (count int) {
	for _, b := range in {
		count += bits.OnesCount8(uint8(b))
	}
	return
}

func buildGrid(s string) (out [][]bool) {
	out = make([][]bool, 0)
	for i := 0; i < 128; i++ {
		cur := fmt.Sprintf("%s-%d", s, i)
		col := make([]bool, 0)
		for _, c := range fmt.Sprintf("%08b", knotByte([]byte(cur))) {
			switch c {
			case '1':
				col = append(col, true)
			case '0':
				col = append(col, false)
			}
		}
		out = append(out, col)
	}
	return
}

// thanks to https://www.reddit.com/user/u794575248
func countRegions(m [][]bool) (res int) {
	seen := make(map[[2]int]bool)

	for i, row := range m {
		for j, bit := range row {
			if bit && !seen[[2]int{i, j}] {
				res++
				q := make([][2]int, 1)
				q[0] = [2]int{i, j}
				for len(q) > 0 {
					seen[q[0]] = true
					x, y := q[0][0], q[0][1]
					q = q[1:]
					for _, p := range [][2]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}} {
						x2, y2 := p[0], p[1]
						if x2 < 128 && x2 >= 0 &&
							y2 < 128 && y2 >= 0 &&
							m[x2][y2] &&
							!seen[[2]int{x2, y2}] {
							q = append(q, [2]int{x2, y2})
						}
					}
				}
			}
		}
	}
	return
}
