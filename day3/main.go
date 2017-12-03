package main

import (
	"fmt"
	"math"
)

func main() {
	input := 361527
	fmt.Printf("First puzzle result: %d \n", getDistance(input))
	fmt.Printf("Second puzzle result: %d \n", createSpiralStress(input))
}

type tuple struct {
	x int
	y int
}

func createSpiral(n int) [][]int {
	//                Right, Up, Left, Down
	var dir = []tuple{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	size := int(math.Sqrt(float64(n)))
	field := make([][]int, size)

	for i := range field {
		field[i] = make([]int, size)
	}

	x, y := size/2, size/2
	if size%2 == 0 {
		y = size/2 - 1
	}
	cur := 1
	steps, stepsInDir, stepsInDirCount := 0, 1, 0
	for i := 0; i < n; i++ {
		field[x][y] = i + 1
		if steps == stepsInDir {
			steps = 0
			stepsInDirCount++

			if stepsInDirCount == 2 {
				stepsInDir++
				stepsInDirCount = 0
			}
			cur = (cur + 1) % 4
		}
		x += dir[cur].x
		y += dir[cur].y
		steps++
	}
	return field
}

func createSpiralStress(n int) int {
	//                Right, Up, Left, Down
	var dir = []tuple{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	size := int(math.Sqrt(float64(n)))
	field := make([][]int, size)

	for i := range field {
		field[i] = make([]int, size)
	}

	x, y := size/2, size/2
	if size%2 == 0 {
		y = size/2 - 1
	}
	cur := 1
	steps, stepsInDir, stepsInDirCount := 0, 1, 0
	for i := 0; i < n; i++ {
		if i == 0 {
			field[x][y] = i + 1
		} else {
			field[x][y] = adjSum(field, x, y)
		}

		if field[x][y] > n {
			return field[x][y]
		}

		if steps == stepsInDir {
			steps = 0
			stepsInDirCount++

			if stepsInDirCount == 2 {
				stepsInDir++
				stepsInDirCount = 0
			}
			cur = (cur + 1) % 4
		}
		x += dir[cur].x
		y += dir[cur].y
		steps++
	}
	return 0
}

//not the most elegant solution :-(
func adjSum(field [][]int, x, y int) (sum int) {
	maxIndex := len(field[0]) - 1
	if x+1 <= maxIndex {
		sum += field[x+1][y]
	}
	if y+1 <= maxIndex {
		sum += field[x][y+1]
	}
	if x+1 <= maxIndex && y+1 <= maxIndex {
		sum += field[x+1][y+1]
	}
	if x-1 >= 0 {
		sum += field[x-1][y]
	}
	if x-1 >= 0 && y+1 <= maxIndex {
		sum += field[x-1][y+1]
	}
	if y-1 >= 0 {
		sum += field[x][y-1]
	}
	if y-1 >= 0 && x-1 >= 0 {
		sum += field[x-1][y-1]
	}
	if y-1 >= 0 && x+1 <= maxIndex {
		sum += field[x+1][y-1]
	}
	return sum
}

func findElement(field [][]int, e int) (x, y int, found bool) {
	for y, row := range field {
		for x, v := range row {
			if e == v {
				return x, y, true
			}
		}
	}
	return 0, 0, false
}

func getDistance(n int) int {
	r := int(math.Sqrt(float64(n))) + 1
	s := createSpiral(r * r)
	x0, y0, _ := findElement(s, 1)
	x1, y1, _ := findElement(s, n)
	x := x0 - x1
	y := y0 - y1
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

func printSpiral(spiral [][]int) {
	for _, row := range spiral {
		for _, v := range row {
			fmt.Printf(" %3d ", v)
		}
		fmt.Print("\n")
	}
}
