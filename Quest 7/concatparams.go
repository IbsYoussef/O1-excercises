package piscine

func ConcatParams(args []string) string {
	var concat string
	for i := 0; i < len(args); i++ {
		if i < len(args)-1 {
			concat += args[i] + "\n"
		} else {
			concat += args[i]
		}
	}
	return concat
}
