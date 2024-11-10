package main

import (
	"fmt"
	"github.com/hacknyack/learning-go/introduction/calculator"
)

func main() {
	fmt.Println("5 + 10 = ", calculator.Add(5, 10))
	fmt.Println(calculator.Divide(18, 10))
	fmt.Println(calculator.Sum(0, 10))
	fmt.Println(calculator.SumUntil(0))
	fmt.Println(calculator.Abs(5))
	fmt.Println(calculator.IsSquareNumber(25))
	fmt.Println(calculator.SumFromAToB(1, 4))
	fmt.Println(calculator.MultiplyFromAToB(1, 10))
}
