package pile

func Rb(pile *Pile) {
	if len(pile.b) < 2 {
		return
	}
	pile.b = append(pile.b[1:], pile.b[0])
}
