package main

import "fmt"

type distance struct {
	from string
	to   string
	d    int
}

func permutations(iterable []int, r int) {
	pool := iterable
	n := len(pool)

	if r > n {
		return
	}

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	result := make([]int, r)
	for i, el := range indices[:r] {
		result[i] = pool[el]
	}

	fmt.Println(result)

	for n > 0 {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j++ {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]

				for k := i; k < r; k++ {
					result[k] = pool[indices[k]]
				}

				fmt.Println(result)

				break
			}
		}

		if i < 0 {
			return
		}

	}

}

func main() {
	distances := []distance{}

	distances = append(distances, distance{from: "London", to: "Dublin", d: 464})
	distances = append(distances, distance{from: "London", to: "Belfast", d: 518})
	distances = append(distances, distance{from: "Dublin", to: "Belfast", d: 141})

	locations := []string{"Belfast", "Dublin", "London"}

	fmt.Println(distances)
	fmt.Println(locations)
	permutations([]int{1, 2, 3}, 3)
}
