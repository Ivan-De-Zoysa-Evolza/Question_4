package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		num1 := getValidDecimalInput("\nEnter the first number (up to 2 decimal places): ")
		operator := getValidOperator()
		num2 := getValidDecimalInput("\nEnter the second number (up to 2 decimal places): ")

		result := calculate(num1, num2, operator)
		fmt.Printf("\nResult: %.2f\n", result)

		fmt.Print("\nDo you want to continue? (yes/no): ")
		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) != "yes" {
			fmt.Println("\nGoodbye!")
			break
		}
	}
}

func getValidDecimalInput(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)

		// Check if it's a valid float
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("\nInvalid number. Please enter a numeric value.")
			continue
		}

		// Allow integers or floats with max 2 decimal places
		if strings.Contains(input, ".") {
			parts := strings.Split(input, ".")
			if len(parts[1]) > 2 {
				fmt.Println("\nOnly up to 2 decimal places allowed.")
				continue
			}
		}

		return value
	}
}

func getValidOperator() string {
	for {
		fmt.Print("Enter the operator (+, -, *, /): ")
		var op string
		fmt.Scanln(&op)
		if op == "+" || op == "-" || op == "*" || op == "/" {
			return op
		}
		fmt.Println("Invalid operator. Please use one of +, -, *, /")
	}
}

func calculate(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		} else {
			fmt.Println("Error: Division by zero.")
			return 0
		}
	default:
		fmt.Println("Unexpected error.")
		return 0
	}
}
