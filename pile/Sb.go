package pile

func Sb(pile *Pile) {
	if len(pile.b) < 2 {
		return
	}
	pile.b[0], pile.b[1] = pile.b[1], pile.b[0]
}
