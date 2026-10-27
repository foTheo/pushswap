package pile

func Rra(pile *Pile) {
	if len(pile.a) < 2 {
		return
	}
	last := pile.a[len(pile.a)-1]
	pile.a = append([]int{last}, pile.a[:len(pile.a)-1]...)
}
