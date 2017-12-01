package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	input, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	result := antiCaptcha(string(input))
	fmt.Printf("First puzzle result: %d \n", result)

	result = antiCaptchaHalfway(string(input))
	fmt.Printf("Second puzzle result: %d \n", result)
}

func antiCaptcha(in string) (sum int) {
	for i := 0; i < len(in); i++ {
		next := (i + 1) % len(in)
		if in[i] == in[next] {
			sum += int(in[i] - '0')
		}
	}
	return
}

func antiCaptchaHalfway(in string) (sum int) {
	for i := 0; i < len(in); i++ {
		next := (i + len(in)/2) % len(in)
		if in[i] == in[next] {
			sum += int(in[i] - '0')
		}
	}
	return
}
