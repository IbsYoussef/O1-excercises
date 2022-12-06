func UltimateDivMod(a *int, b *int) {
	inta := *a

	*a = inta / *b
	*b = inta % *b

}