package days

import (
	"bufio"
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type ingredient struct {
	capacity   int
	durability int
	frosting   int
	texture    int
	calories   int
}

type amounts struct {
	a int
	b int
	c int
	d int
}

// Generate ingredient amount whose sum is cap
func ingredientAmounts(start, cap int) <-chan amounts {
	ch := make(chan (amounts))
	go func() {
		for a := start; a < cap; a++ {
			for b := start; b < cap; b++ {
				if a+b > cap {
					continue
				}
				for c := start; c < cap; c++ {
					if a+b+c > cap {
						continue
					}
					for d := start; d < cap; d++ {
						if (a + b + c + d) == cap {
							ch <- amounts{a, b, c, d}
						}
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}

// Run15 will help us make the best cookies
func Run15(scanner *bufio.Scanner) string {
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	maxScore := 0
	calorieScore := 0
	sprinkles := ingredient{capacity: 5, durability: -1, frosting: 0, texture: 0, calories: 5}
	peanutButter := ingredient{capacity: -1, durability: 3, frosting: 0, texture: 0, calories: 1}
	frosting := ingredient{capacity: 0, durability: -1, frosting: 4, texture: 0, calories: 6}
	sugar := ingredient{capacity: -1, durability: 0, frosting: 0, texture: 2, calories: 8}

	for i := range ingredientAmounts(1, 100) {
		score :=
			max(0, sprinkles.capacity*i.a+peanutButter.capacity*i.b+frosting.capacity*i.c+sugar.capacity*i.d) *
				max(0, sprinkles.durability*i.a+peanutButter.durability*i.b+frosting.durability*i.c+sugar.durability*i.d) *
				max(0, sprinkles.frosting*i.a+peanutButter.frosting*i.b+frosting.frosting*i.c+sugar.frosting*i.d) *
				max(0, sprinkles.texture*i.a+peanutButter.texture*i.b+frosting.texture*i.c+sugar.texture*i.d)

		if score > maxScore {
			maxScore = score
		}

		if (sprinkles.calories*i.a+peanutButter.calories*i.b+frosting.calories*i.c+sugar.calories*i.d) == 500 &&
			score > calorieScore {
			calorieScore = score
		}
	}
	return fmt.Sprintf("%v, %d, %v, %d", maxScore == 13882464, maxScore, calorieScore == 11171160, calorieScore)
}
