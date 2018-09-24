package days

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/madboy/adventofcode/2015/tools"
)

const nbrSpells = 5

type player struct {
	hp      int
	mana    int
	armor   int
	dmg     int
	effects [5]spell
	timers  []int
	cost    int
}

func (p player) String() string {
	return fmt.Sprintf("{hp %d, armor: %d, mana: %d,}",
		p.hp, p.armor, p.mana,
	)
}

type boss struct {
	hp  int
	dmg int
}

func (b boss) String() string {
	return fmt.Sprintf("{hp: %d, dmg: %d}", b.hp, b.dmg)
}

type spell struct {
	name   string
	cost   int
	hp     int
	dmg    int
	armor  int
	timer  int
	effect func(p *player)
	idx    int
}

func (s spell) String() string {
	return fmt.Sprintf("timer: %d, cost: %d, hp: %d, dmg: %d",
		s.timer, s.cost, s.hp, s.dmg,
	)
}
func poisonE(p *player) {
	p.dmg += 3
}

func rechargeE(p *player) {
	p.mana += 101
}

func shieldE(p *player) {
	p.armor += 7
}

func (p *player) applyEffects() {
	for i, f := range p.effects {
		if f.effect != nil && p.timers[i] > 0 {
			f.effect(p)
		}
	}
}

func (p player) updateTimers() {
	for i := range p.effects {
		if p.timers[i] > 0 {
			p.timers[i]--
		}
	}
}

func applySpell(p *player, s spell) {
	p.dmg += s.dmg
	p.hp += s.hp
	p.armor += s.armor
}

func magicFight(p *player, b *boss, pTurn bool, s spell) bool {
	p.applyEffects()
	p.updateTimers()
	if pTurn && p.hp > 0 {
		applySpell(p, s)
		if p.mana < 0 {
			fmt.Println("oom")
			return false
		}
		p.mana -= s.cost
		p.cost += s.cost
		b.hp -= p.dmg
		if s.effect != nil {
			p.effects[s.idx] = s
			p.timers[s.idx] = s.timer
		}
	} else {
		dmg := b.dmg - p.armor
		b.hp -= p.dmg
		if b.hp > 0 {
			p.hp -= tools.Max(dmg, 1)
		}
	}
	p.dmg = 0
	p.armor = 0
	return b.hp <= 0 && p.hp > 0
}

func minCost(arr []int) int {
	m := math.MaxInt32
	for _, v := range arr {
		if v < m {
			m = v
		}
	}
	return m
}

func runFight(spells []spell, difficulty string) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var effects [nbrSpells]spell
	timers := []int{0, 0, 0, 0, 0}
	p := player{hp: 50, mana: 500, dmg: 0, armor: 0, effects: effects, timers: timers}
	b := boss{58, 9}
	playerWon := false

	for p.hp > 0 && p.mana > 0 && !playerWon {
		var s spell
		for i := 0; i < 10; i++ {
			s = spells[r1.Intn(nbrSpells)]
			if s.cost < p.mana && p.timers[s.idx] <= 1 {
				break
			}
		}
		if p.timers[s.idx] > 1 {
			break
		}

		if difficulty == "hard" {
			p.hp--
		}
		playerWon = magicFight(&p, &b, true, s)

		if !playerWon {
			playerWon = magicFight(&p, &b, false, s)
		}
	}
	if playerWon {
		return p.cost
	}
	return -1
}

// Run22 in which we are fighting the boss with magic
func Run22(scanner *bufio.Scanner) string {
	poison := spell{name: "poison", cost: 173, timer: 6, effect: poisonE, idx: 0}
	mm := spell{name: "mm", cost: 53, dmg: 4, idx: 1}
	recharge := spell{name: "recharge", cost: 229, timer: 5, effect: rechargeE, idx: 2}
	shield := spell{name: "shield", cost: 113, timer: 6, effect: shieldE, idx: 3}
	drain := spell{name: "drain", cost: 73, dmg: 2, hp: 2, idx: 4}
	spells := []spell{recharge, shield, drain, poison, mm}

	var normalCosts []int
	var hardCosts []int
	// 100000 random iterations seems to yield consistent results
	for t := 0; t < 100000; t++ {
		result := runFight(spells, "normal")
		if result > 0 {
			normalCosts = append(normalCosts, result)
		}

		result = runFight(spells, "hard")
		if result > 0 {
			hardCosts = append(hardCosts, result)
		}
	}
	// part1: 1269, part2: 1309
	return fmt.Sprintf("Least amount of mana I can spend on normal and win is: %d\nLeast amoutn of mana I can spend on hard and win is: %d", minCost(normalCosts), minCost(hardCosts))
}
