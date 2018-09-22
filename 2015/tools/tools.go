package tools

import "math"

// Range return a list of integers from low .. high-1
func Range(low, high int) []int {
	var r []int
	for i := low; i < high; i++ {
		r = append(r, i)
	}
	return r
}

// Sum all elements in an array
func Sum(arr []float64) float64 {
	sum := 0.0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// GetDivisors calculates the diviors of a number
func GetDivisors(n float64) []float64 {
	var divisors []float64
	for i := 1.0; i <= math.Sqrt(n); i++ {
		if math.Mod(n, i) == 0 {
			if n/i == i {
				divisors = append(divisors, i)
			} else {
				divisors = append(divisors, i)
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}

// Max of x and y
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
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
