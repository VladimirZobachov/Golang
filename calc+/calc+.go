package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter a mathematical expression: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	result := calculate(input)
	fmt.Println("Input:", input)
	fmt.Println("Output:", result)
}

func calculate(input string) string {
	var result float64
	operations := strings.Split(input, " ")

	result, _ = strconv.ParseFloat(operations[0], 64)
	for i := 1; i < len(operations); i += 2 {
		operator := operations[i]
		value := strings.Trim(operations[i+1], "\n")
		parsedValue, _ := strconv.ParseFloat(value, 64)

		switch operator {
		case "+":
			result += parsedValue
		case "-":
			result -= parsedValue
		case "*":
			result *= parsedValue
		case "/":
			result /= parsedValue
		}
	}

	return fmt.Sprintf("%.2f", result)
}
