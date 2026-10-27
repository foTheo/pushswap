package pile

func Pb(pile *Pile) {
	if len(pile.a) == 0 {
		return
	}
	pile.b = append([]int{pile.a[0]}, pile.b...)
	pile.a = pile.a[1:]
}
