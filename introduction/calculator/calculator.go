package calculator

// package level
var Pi = 3.141592654

func Add(left int, right int) int {
	//var sum int = left + right

	//var sum = left + right

	//var sum int
	//sum = left + right

	sum := left + right

	return sum
}

func Sub(left, right int) int {
	return left - right
}

func Divide(left, right int) (quotient int, remainder int) {
	quotient = left / right
	remainder = left % right

	// naked return
	return
}
