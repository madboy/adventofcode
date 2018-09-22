package days

import (
	"bufio"
	"fmt"

	"github.com/madboy/adventofcode/2015/tools"
	"github.com/madboy/combinatoric"
)

// Item representation
type Item struct {
	cost  int
	dmg   int
	armor int
}

// Creature is for both player and boss
type Creature struct {
	hp    int
	dmg   int
	armor int
}

func fight(p, b Creature) bool {
	pTurn := true
	for p.hp > 0 && b.hp > 0 {
		if pTurn {
			dmg := p.dmg - b.armor
			b.hp -= tools.Max(dmg, 1)
		} else {
			dmg := b.dmg - p.armor
			p.hp -= tools.Max(dmg, 1)
		}
		pTurn = !pTurn
	}

	return p.hp > b.hp
}

// Run21 in which we are counting presents
func Run21(scanner *bufio.Scanner) string {
	weapons := []Item{
		Item{8, 4, 0},
		Item{10, 5, 0},
		Item{25, 6, 0},
		Item{40, 7, 0},
		Item{74, 8, 0},
	}

	armor := []Item{
		Item{13, 0, 1},
		Item{31, 0, 2},
		Item{53, 0, 3},
		Item{75, 0, 4},
		Item{102, 0, 5},
	}

	rings := []Item{
		Item{25, 1, 0},
		Item{50, 2, 0},
		Item{100, 3, 0},
		Item{20, 0, 1},
		Item{40, 0, 2},
		Item{80, 0, 3},
	}

	minCost := 10000
	maxCost := 0
	for _, w := range weapons {
		for _, a := range tools.Range(-1, 5) {
			for _, r := range combinatoric.Combinations(tools.Range(-2, 6), 2) {
				boss := Creature{109, 8, 2}
				player := Creature{100, 0, 0}
				cost := 0

				player.dmg += w.dmg
				cost += w.cost

				// Armor is optional
				if a != -1 {
					player.armor += armor[a].armor
					cost += armor[a].cost
				}

				// We can have 0, 1 or 2 rings on
				r0 := r[0]
				r1 := r[1]

				if r0 > 0 {
					player.dmg += rings[r0].dmg
					player.armor += rings[r0].armor
					cost += rings[r0].cost
				}

				if r1 > 0 {
					player.dmg += rings[r1].dmg
					player.armor += rings[r1].armor
					cost += rings[r1].cost
				}

				won := fight(player, boss)
				if won && cost < minCost {
					minCost = cost
				}

				if !won && cost > maxCost {
					maxCost = cost
				}
			}
		}
	}

	return fmt.Sprintf("Least amount of gold I can spend and win is: %d\nMost amount of gold I can spend and still loose is: %d", minCost, maxCost)
}
