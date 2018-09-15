package tools

// Range return a list of integers from low .. high-1
func Range(low, high int) []int {
	var r []int
	for i := low; i < high; i++ {
		r = append(r, i)
	}
	return r
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
