package main

import (
	"fmt"
	"os"
	"strconv"
)

const mod = 2147483647

func main() {
	args := os.Args
	var a, b uint64 = 65, 8921
	if len(args) != 3 {
		fmt.Printf("Use: %s a b\n", args[0])
		fmt.Printf("No Input was given will, use %d %d instead\n\n", a, b)
	} else {
		a1, err := strconv.Atoi(args[1])
		b1, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err, "\nError: a and b must be integer values")
			os.Exit(2)
		}
		a = uint64(a1)
		b = uint64(b1)
	}
	fmt.Printf("The final count after 40 Million pairs is %d.\n", partOne(a, b, 4E7))
	fmt.Printf("The final count after  5 Million pairs is %d.\n", partTwo(a, b, 5E6))
}
func partOne(a, b uint64, pairsNb int) (res int) {

	for i := 0; i < pairsNb; i++ {
		a *= 16807
		a %= mod
		b *= 48271
		b %= mod

		if (a & 0xffff) == (b & 0xffff) {
			res++
		}
	}
	return res
}

func partTwo(a, b uint64, pairsNb int) (res int) {

	for i := 0; i < pairsNb; i++ {
		for {
			a *= 16807
			a %= mod
			if a%4 == 0 {
				break
			}
		}
		for {
			b *= 48271
			b %= mod
			if b%8 == 0 {
				break
			}
		}
		if (a & 0xffff) == (b & 0xffff) {
			res++
		}
	}
	return res
}
