package calculator

func Sum(min, max int) int {
	sum := 0
	// for with condition
	for i := min; i < max; i++ {
		sum += i
	}
	return sum
}

func SumUntil(max int) int {
	sum := 1
	// while
	for sum < max {
		sum += sum
	}
	return sum
}

func SumInfinite() {
	for {
		//...
	}
}
