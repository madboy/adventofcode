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
	spoons int
	c      int
	d      int
	f      int
	t      int
	j      int
}

func (i ingredient) capacity() int {
	return i.spoons * i.c
}

func (i ingredient) durability() int {
	return i.spoons * i.d
}

func (i ingredient) flavor() int {
	return i.spoons * i.f
}

func (i ingredient) texture() int {
	return i.spoons * i.t
}

func (i ingredient) calories() int {
	return i.spoons * i.j
}

// Run15 will help us make the best cookies
func Run15(scanner *bufio.Scanner) string {
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	maxScore := 0
	calorieScore := 0
	sprinkles := ingredient{c: 5, d: -1, f: 0, t: 0, j: 5}
	peanutButter := ingredient{c: -1, d: 3, f: 0, t: 0, j: 1}
	frosting := ingredient{c: 0, d: -1, f: 4, t: 0, j: 6}
	sugar := ingredient{c: -1, d: 0, f: 0, t: 2, j: 8}

	for s := 1; s < 100; s++ {
		sprinkles.spoons = s
		for p := 1; p < 100; p++ {
			peanutButter.spoons = p
			for f := 1; f < 100; f++ {
				frosting.spoons = f
				for s := 1; s < 100; s++ {
					sugar.spoons = s
					if (sprinkles.spoons + peanutButter.spoons + frosting.spoons + sugar.spoons) == 100 {
						score :=
							max(0, sprinkles.capacity()+peanutButter.capacity()+frosting.capacity()+sugar.capacity()) *
								max(0, sprinkles.durability()+peanutButter.durability()+frosting.durability()+sugar.durability()) *
								max(0, sprinkles.flavor()+peanutButter.flavor()+frosting.flavor()+sugar.flavor()) *
								max(0, sprinkles.texture()+peanutButter.texture()+frosting.texture()+sugar.texture())

						if score > maxScore {
							maxScore = score
						}

						if (sprinkles.calories()+peanutButter.calories()+frosting.calories()+sugar.calories()) == 500 &&
							score > calorieScore {
							calorieScore = score
						}
					}

				}
			}
		}
	}
	return fmt.Sprintf("%v, %d, %d", maxScore == 13882464, maxScore, calorieScore)
}
