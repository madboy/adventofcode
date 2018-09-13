package days

import (
	"bufio"
	"fmt"
)

// Reindeer describe the reindeer properties we are interested in
type Reindeer struct {
	Name      string
	Speed     int
	Rest      int
	Fly       int
	Points    int
	CycleTime int
	Distance  int
}

func (r *Reindeer) travel() int {
	if r.CycleTime < r.Fly {
		r.Distance += r.Speed
	}

	if r.CycleTime < (r.Fly + r.Rest) {
		r.CycleTime++
	}

	if r.CycleTime%(r.Fly+r.Rest) == 0 {
		r.CycleTime = 0
	}
	return r.Distance
}

// Run14 will tell us the travel distance of the winning reindeer
func Run14(scanner *bufio.Scanner) string {
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var reindeers []*Reindeer

	reindeers = append(reindeers, &Reindeer{Name: "Dancer", Speed: 27, Fly: 5, Rest: 132})
	reindeers = append(reindeers, &Reindeer{Name: "Cupid", Speed: 22, Fly: 2, Rest: 41})
	reindeers = append(reindeers, &Reindeer{Name: "Rudolph", Speed: 11, Fly: 5, Rest: 48})
	reindeers = append(reindeers, &Reindeer{Name: "Donner", Speed: 28, Fly: 5, Rest: 134})
	reindeers = append(reindeers, &Reindeer{Name: "Dasher", Speed: 4, Fly: 16, Rest: 55})
	reindeers = append(reindeers, &Reindeer{Name: "Blitzen", Speed: 14, Fly: 3, Rest: 38})
	reindeers = append(reindeers, &Reindeer{Name: "Prancer", Speed: 3, Fly: 21, Rest: 40})
	reindeers = append(reindeers, &Reindeer{Name: "Comet", Speed: 18, Fly: 6, Rest: 103})
	reindeers = append(reindeers, &Reindeer{Name: "Vixen", Speed: 18, Fly: 5, Rest: 84})

	// reindeers = append(reindeers, &Reindeer{Name: "Comet", Speed: 14, Fly: 10, Rest: 127})
	// reindeers = append(reindeers, &Reindeer{Name: "Dancer", Speed: 16, Fly: 11, Rest: 162})
	maxDistance := 0
	maxPoints := 0

	for t := 0; t < 2503; t++ {
		for _, r := range reindeers {
			distance := r.travel()
			if distance > maxDistance {
				maxDistance = distance
			}
		}

		// We don't know if we are in the lead and should get points until all the reindeer have travelled
		for _, r := range reindeers {
			if r.Distance == maxDistance {
				r.Points++
			}

			if r.Points > maxPoints {
				maxPoints = r.Points
			}
		}
	}

	return fmt.Sprintf("The winning reindeer on distance traveled: %d\nWhile the winning reindeer on points got: %d", maxDistance, maxPoints)
}
