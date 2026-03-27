package main

import (
	"fmt"
)

func main() {

	for {
		var help string
		var a, b float64
		var op string
		fmt.Println(help, "For the help in this calculator use + to add, - to substract, * to multiply, / to divide.")

		fmt.Print("Enter first number: ")
		fmt.Scan(&a)

		fmt.Print("Enter operator (+, -, *, /): ")
		fmt.Scan(&op)

		fmt.Print("Enter second number: ")
		fmt.Scan(&b)

		switch op {
		case "+":
			fmt.Println("Answer:", a+b)
		case "-":
			fmt.Println("Answer:", a-b)
		case "*":
			fmt.Println("Answer:", a*b)
		case "/":
			if b == 0 {
				fmt.Println("Can't divide by zero!")
			} else {
				fmt.Println("Answer:", a/b)
			}
		default:
			fmt.Println("Invalid operator!")
		}
	}
}
