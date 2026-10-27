package pile

func Ra(p *Pile) {
	if len(p.a) < 2 {
		return
	}
	p.a = append(p.a[1:], p.a[0])
}
