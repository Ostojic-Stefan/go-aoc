package utils

type Pair struct {
	X int
	Y int
}

func (p *Pair) Add(other Pair) Pair {
	return Pair{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}
