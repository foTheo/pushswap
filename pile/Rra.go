package pile

func Rra(pile *Pile) {
	if len(pile.a) < 2 {
		return
	}
	dernier := pile.a[len(pile.a)-1]
	pile.a = append([]int{dernier}, pile.a[:len(pile.a)-1]...)
}
