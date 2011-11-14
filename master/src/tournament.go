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
