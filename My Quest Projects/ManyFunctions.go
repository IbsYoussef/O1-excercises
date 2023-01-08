func Function_1(nbr int) int {
	result := nbr * 100 / 2
	return result
}

func Function_2(nbr int) int {
	result := nbr / 10 * 50
	return result
}

func Function_3(nbr int) int {
	result := 0
	for i := 0; i < nbr; i++ {
		result = result + i
	}
	return result
}

func ApplyFunction(f func(int) int, answer int) int {
	answer = f(answer)
	return answer
}

func main() {
	ArrayFunc := []func(int) int{Function_1, Function_2, Function_3}

	fmt.Println("Choose a number to apply a function to")
	var n int
	fmt.Scanln(&n)

	fmt.Println("Choose the function you want to apply, Options of functions are from 0 - 2")
	var x int
	fmt.Scanln(&x)

	fmt.Println("Final Result:")
	fmt.Println(ApplyFunction(ArrayFunc[x], n))

}
