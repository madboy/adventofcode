package tools

import "math"

//Permutations return successive r length permutations of elements in values.
//Using Heap's algorithm, https://en.wikipedia.org/wiki/Heap%27s_algorithm
func Permutations(values []int) [][]int {
	var c []int
	var permutations [][]int
	n := len(values)

	for i := 0; i < n; i++ {
		c = append(c, 0)
	}

	tmp := make([]int, len(values))
	copy(tmp, values)
	permutations = append(permutations, tmp)

	for i := 0; i < n; {
		if c[i] < i {
			if i%2 == 0 {
				swap(values, 0, i)
			} else {
				swap(values, c[i], i)
			}
			tmp := make([]int, len(values))
			copy(tmp, values)
			permutations = append(permutations, tmp)
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	return permutations
}

// Combinations retruns unique combinations of values of length r
// Implementation from python itertools https://docs.python.org/3.6/library/itertools.html#itertools.combinations
func Combinations(values []int, r int) [][]int {
	var combinations [][]int
	pool := make([]int, len(values))
	copy(pool, values)
	n := len(pool)
	if r > n {
		return combinations
	}
	indices := Range(r)

	var combination []int
	for _, i := range indices {
		combination = append(combination, pool[i])
	}
	combinations = append(combinations, combination)

	for {
		ii := math.MinInt32
		for _, i := range Reversed(Range(r)) {
			if indices[i] != i+n-r {
				ii = i
				break
			}
		}
		if ii == math.MinInt32 {
			return combinations
		}
		indices[ii]++
		for _, j := range XRange(ii+1, r) {
			indices[j] = indices[j-1] + 1
		}
		combinations = append(combinations, createCombination(indices, pool))

	}
	return combinations
}

func createCombination(indices, values []int) []int {
	var combination []int
	for _, i := range indices {
		combination = append(combination, values[i])
	}
	return combination
}

func swap(values []int, f, t int) {
	values[f], values[t] = values[t], values[f]
}

// Range return a list of integers from 0 .. n-1
func Range(n int) []int {
	var r []int
	for i := 0; i < n; i++ {
		r = append(r, i)
	}
	return r
}

// XRange returns a list of integers from low .. high -1
func XRange(low, high int) []int {
	var r []int
	for i := low; i < high; i++ {
		r = append(r, i)
	}
	return r
}

// Reversed returns a list of integers in reverse order
func Reversed(values []int) []int {
	var reversed []int
	for i := len(values) - 1; i >= 0; i-- {
		reversed = append(reversed, values[i])
	}
	return reversed
}

// Set is an unordered collection of unique int elements
type Set struct {
	Values []string
	keys   map[string]bool
}

// NewSet returns an empty set
func NewSet() Set {
	return Set{keys: make(map[string]bool)}
}

// Add value to the set
func (s *Set) Add(value string) {
	if !s.keys[value] {
		s.keys[value] = true
		s.Values = append(s.Values, value)
	}
}
