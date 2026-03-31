package main

import (
	"fmt"
)

func main() {
	for {
		var input string
		var a, b float64
		var op string
		fmt.Println("\nEnter a number, 'help' for instructions, or 'exit' to quit:")
		fmt.Print("> ")
		fmt.Scan(&input)
		if input == "help" {
			showHelp()
			continue
		}
		if input == "exit" {
			fmt.Println("Thank you for using chekus calculator. Goodbye!")
			break
		}
		_, err1 := fmt.Sscanf(input, "%f", &a)
		if err1 != nil {
			fmt.Println("Error: That's not a number. Try again.")
			continue
		}
		fmt.Print("Enter operator (+, -, *, /): ")
		fmt.Scan(&op)
		fmt.Print("Enter second number: ")
		_, err2 := fmt.Scan(&b)
		if err2 != nil {
			fmt.Println("Error: Invalid number. Try again.")
			var trash string
			fmt.Scanln(&trash)
			continue
		}
		switch op {
		case "+":
			fmt.Println("Answer:", a+b)
		case "-":
			fmt.Println("Answer:", a-b)
		case "*":
			fmt.Println("Answer:", a*b)
		case "/":
			if b == 0 {
				fmt.Println("Error: Can't divide by zero!")
			} else {
				fmt.Println("Answer:", a/b)
			}
		default:
			fmt.Println("Invalid operator!")
		}
	}
}
func showHelp() {
	fmt.Println("\nInstructions")
	fmt.Println("1. addition (+)")
	fmt.Println("2. subtract (-)")
	fmt.Println("3. multiplication (*)")
	fmt.Println("4. division (/)")
	fmt.Println("Type 'exit' to quit.")
}
