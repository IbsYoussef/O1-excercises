package piscine

func AppendRange(min, max int) []int {
	array := []int{}
	if min >= max {
		array = nil
	} else {
		for ; min < max; min++ {
			array = append(array, min)
		}
	}
	return array
}
