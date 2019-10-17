package piscine

func UltimateDivMod(a *int, b *int) {
	divid := *a / *b
	rest := *a
	*a = divid
	modULO := rest % *b
	*b = modULO

}
