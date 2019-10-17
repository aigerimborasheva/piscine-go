package piscine

func DivMod(a int, b int, div *int, mod *int) {
	divid := a/b
	*div = divid 
	modULO := a%b
	*mod = modULO
}


