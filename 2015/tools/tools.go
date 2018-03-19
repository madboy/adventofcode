package tools

// Permutations return successive r length permutations of elements in values.
// Using Heap's algorithm, https://en.wikipedia.org/wiki/Heap%27s_algorithm
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
