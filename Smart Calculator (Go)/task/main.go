package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func showHelp() {
	fmt.Println("The program calculates teh sum and difference of numbers.")
	fmt.Println("It supports multiple operators (+ and -), including unary minus and repeated operators.")
}

func normalizeOperators(expression string) string {
	re := regexp.MustCompile(`[-+]+`)
	return re.ReplaceAllStringFunc(expression, func(op string) string {
		negatives := strings.Count(op, "-")
		if negatives%2 == 0 {
			return "+"
		}
		return "-"
	})
}

func calculateExpression(expression string) (int, error) {
	expression = normalizeOperators(expression)
	parts := strings.Fields(expression)
	if len(parts) == 0 {
		return 0, fmt.Errorf("invalid input")
	}

	result, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid input")
	}

	for i := 1; i < len(parts); i += 2 {
		if i+1 >= len(parts) {
			return 0, fmt.Errorf("invalid input")
		}
		operator := parts[i]
		num, err := strconv.Atoi(parts[i+1])
		if err != nil {
			return 0, fmt.Errorf("invalid input")
		}
		switch operator {
		case "+":
			result += num
		case "-":
			result -= num
		default:
			return 0, fmt.Errorf("invalid input")
		}
	}
	return result, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "/exit" {
			fmt.Println("Bye!")
			break
		}

		if input == "/help" {
			showHelp()
			continue
		}

		if input == "" {
			continue
		}

		result, err := calculateExpression(input)
		if err != nil {
			fmt.Println("Invalid input")
		} else {
			fmt.Println(result)
		}
	}
}
