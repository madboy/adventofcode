package tools

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input: []int{0, 1, 2},
			expected: [][]int{
				[]int{0, 1, 2},
				[]int{0, 2, 1},
				[]int{1, 0, 2},
				[]int{1, 2, 0},
				[]int{2, 0, 1},
				[]int{2, 1, 0},
			},
		},
		{
			input: []int{0, 1, 2, 3},
			expected: [][]int{
				[]int{0, 1, 2, 3},
				[]int{0, 1, 3, 2},
				[]int{0, 2, 1, 3},
				[]int{0, 2, 3, 1},
				[]int{0, 3, 1, 2},
				[]int{0, 3, 2, 1},
				[]int{1, 0, 2, 3},
				[]int{1, 0, 3, 2},
				[]int{1, 2, 0, 3},
				[]int{1, 2, 3, 0},
				[]int{1, 3, 0, 2},
				[]int{1, 3, 2, 0},
				[]int{2, 0, 1, 3},
				[]int{2, 0, 3, 1},
				[]int{2, 1, 0, 3},
				[]int{2, 1, 3, 0},
				[]int{2, 3, 0, 1},
				[]int{2, 3, 1, 0},
				[]int{3, 0, 1, 2},
				[]int{3, 0, 2, 1},
				[]int{3, 1, 0, 2},
				[]int{3, 1, 2, 0},
				[]int{3, 2, 0, 1},
				[]int{3, 2, 1, 0},
			},
		},
	}

	for _, test := range tests {
		got := Permutations(test.input)
		if len(got) != len(test.expected) || !allPresent(got, test.expected) {
			t.Error(
				"input", test.input,
				"expected", test.expected,
				"got", got,
			)
		}

	}
}

func allPresent(got, expected [][]int) bool {
	var matches int
	for _, e := range expected {
		for _, g := range got {
			if cmp.Equal(e, g) {
				matches++
				break
			}
		}
	}
	return matches == len(got)
}
