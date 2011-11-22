package main

import (
	"fmt"
	"rand"
	"sort"
)

type priority struct {
	i, j  Team
	value float64
}

func (p priority) String() string {
	return fmt.Sprintf("{%.2f: %s plays %s}", p.value, Team(p.i), Team(p.j))
}

type SortedPriorities []priority

func (p SortedPriorities) Len() int           { return len(p) }
func (p SortedPriorities) Less(i, j int) bool { return p[i].value < p[j].value }
func (p SortedPriorities) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Gene struct {
	Priorities [teams][2 * teams]float64
}

func (g *Gene) Randomize() {
	for i := range g.Priorities {
		for j := range g.Priorities[i] {
			g.Priorities[i][j] = rand.Float64() * 100
		}
		g.Priorities[i][i] = -1
		g.Priorities[i][i+teams] = -1
	}
}

func (g *Gene) Generate(t *Tournament) bool {
	t.Zero()
	//now we stick values in for the teams based on priorities
	//make a copy of priorities
	s := make(SortedPriorities, (teams-1)*teams*2)
	idx := 0
	for i := range g.Priorities {
		for j, val := range g.Priorities[i] {
			if val >= 0 {
				s[idx] = priority{Team(i), Team(j), val}
				idx++
			}
		}
	}

	sort.Sort(s)
	//now we go through the list and assign games, hoping it works!
	var (
		placed     bool
		considered bool
	)
	for {
		placed = false
		considered = false
		for i := range s {
			if s[i].value < 0 { //game has been placed
				continue
			}
			considered = true
			if t.Place(s[i].i, s[i].j) {
				placed = true
				break
			}
		}
		if !placed {
			break
		}
	}
	return !considered
}