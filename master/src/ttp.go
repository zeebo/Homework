package main

const (
	teams = 6
)

func main() {
	var (
		g Gene
		t Tournament
	)
	g.Randomize()

	g.Generate(&t)
}
