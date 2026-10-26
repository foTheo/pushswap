package pile

type Pile struct {
	a []int
	b []int
}

func Sa(pile *Pile) {
	if len(pile.a) < 2 {
		return
	}
	pile.a[0], pile.a[1] = pile.a[1], pile.a[0]
}
