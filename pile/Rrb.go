package pile

func Rrb(pile *Pile) {
	if len(pile.b) < 2 {
		return
	}
	dernier := pile.b[len(pile.b)-1]
	pile.b = append([]int{dernier}, pile.b[:len(pile.b)-1]...)
}
