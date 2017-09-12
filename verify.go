package main

import (
	"fmt"
)

func main() {

	// NOTE: verify that the results correspond to produced page images (100 ~ 103 are on the second page, etc)
	var tests = []int{0, 1, 2, 3, 10, 11, 12, 13, 100, 101, 102, 103}

	for _, offset := range tests {
		fmt.Printf("%v: %X\n", offset, genCode("foobar", offset))
	}
}
