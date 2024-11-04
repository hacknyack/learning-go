package main

import (
	"fmt"
	"github.com/hacknyack/learning-go/introduction/calculator"
)

func main() {
	fmt.Println("5 + 10 = ", calculator.Add(5, 10))
	fmt.Println(calculator.Divide(18, 10))
}
