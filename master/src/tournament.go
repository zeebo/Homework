package main

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

func (t *Tournament) Place(home, away Team) bool {
	//attempt to place the game Home plays vs Away
	// home is [0, teams)
	// away is [0, 2*teams)

	return true
}
