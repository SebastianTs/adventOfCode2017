package main

import "fmt"

const (
	listSize = 256
	block    = 16
)

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

func knotByte(ls []byte) []byte {

	ls = append(ls, []byte{17, 31, 73, 47, 23}...)

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

func knotString(s string) string {
	return fmt.Sprintf("%x", knotByte([]byte(s)))
}

func sToDense(sparse []byte) (dense []byte) {
	for i := 0; i < block; i++ {
		dense = append(dense, sToDenseBlock(sparse[i*block:]))
	}
	return dense
}

func sToDenseBlock(sparse []byte) (dense byte) {
	var res byte
	for j := 0; j < block; j++ {
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
