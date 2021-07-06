package main

import (
	"math/rand"
	"testing"
)

func Test_quickSort(t *testing.T) {
	const sl = 20
	var tests = make([][sl]int, sl)

	for i := range tests {
		for j := 0; j < sl; j++ {
			tests[i][j] = rand.Int()
		}
	}

	for i := range tests {
		quickSort(tests[i][:], 0, len(tests[i][:])-1)
	}

	for i := range tests {
		if !checkSort(tests[i][:]) {
			t.Errorf("жопа")
		}
	}

}

func checkSort(n []int) bool {
	for i := range n {
		if i+1 < len(n) {
			return true
		}

		if i > n[i+1] {
			return false
		}
	}

	return true
}