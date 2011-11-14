package main

import "fmt"

const (
	teams     = 6
	teamNames = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Team int

func (t Team) String() string {
	if t%2 == 0 {
		return fmt.Sprintf("@%s", teamNames[t])
	}
	return fmt.Sprintf(" %s", teamNames[t])
}

type Gene struct {
	Priorities [teams][teams]float64
}

type Tournament [teams][2 * (teams - 1)]Team

func (t *Tournament) Zero() {
	for i := range t {
		for j := range t[i] {
			t[i][j] = 0
		}
	}
}

func (t *Tournament) Valid() bool {
	return false
}

func (g *Gene) Generate(t *Tournament) {
	t.Zero()
}

func main() {
	fmt.Println("Hello World")
}
