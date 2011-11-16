package main

import "fmt"

const teamNames = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Team int

func (t Team) String() string {
	if t < teams {
		return fmt.Sprintf(" %v", string(teamNames[t]))
	}
	return fmt.Sprintf("@%v", string(teamNames[t-teams]))

}

func (t Team) Home() bool {
	return t < teams
}

func (t Team) Away() bool {
	return t >= teams
}

func (t Team) BaseTeam() Team {
	if t.Away() {
		return Team(t - teams)
	}
	return t
}
