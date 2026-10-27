package pile

func Pa(pile *Pile) {
	if len(pile.b) == 0 {
		return
	}
	pile.a = append([]int{pile.b[0]}, pile.a...)
	pile.b = pile.b[1:]
}
