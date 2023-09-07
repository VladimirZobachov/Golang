package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the first and second numbers with space:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	values := strings.Split(input, " ")

	num1 := strings.Trim(values[0], "\n")
	num2 := strings.Trim(values[1], "\n")

	parsedNum1, _ := strconv.ParseFloat(num1, 64)
	parsedNum2, _ := strconv.ParseFloat(num2, 64)

	var operator string
	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Println("Result:", performOperation(parsedNum1, parsedNum2, operator))
}

func performOperation(num1, num2 float64, operator string) float64 {
	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Error: Invalid operator")
	}

	return result
}
