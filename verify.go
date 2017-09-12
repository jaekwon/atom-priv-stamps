package main

import (
	"fmt"
)

func main() {

	var tests = []int{0, 1, 2, 3, 10, 11, 12, 13, 100, 101, 102, 103}

	for _, offset := range tests {
		fmt.Printf("%v: %X\n", offset, genCode("foobar", offset))
	}
}
