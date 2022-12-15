package piscine

// func main(){}

func Map(f func(int) bool, a []int) []bool {
	var tabBool []bool
	for _, element := range a {
		tabBool = append(tabBool, f(element))
	}
	return tabBool
}
