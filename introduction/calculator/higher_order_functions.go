package calculator

//func SumFromAToB(a, b int) int {
//	if a > b {
//		return 0
//	}
//	return a + SumFromAToB(a+1, b)
//}
//
//func MultiplyFromAToB(a, b int) int {
//	if a > b {
//		return 1
//	}
//	return a * MultiplyFromAToB(a+1, b)
//}

func SumFromAToB(a, b int) int {
	return ProcessFromAToB(a, b, 0, func(x int, y int) int {
		return x + y
	})
}

func MultiplyFromAToB(a, b int) int {
	return ProcessFromAToB(a, b, 1, func(x, y int) int {
		return x * y
	})
}

func AddX(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Func which takes another function as parameter
func ProcessFromAToB(a, b, initValue int, fn func(int, int) int) int {
	increment := AddX(1)
	if a > b {
		return initValue
	}
	return fn(a, ProcessFromAToB(increment(a), b, initValue, fn))
}
