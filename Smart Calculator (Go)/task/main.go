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
	fmt.Println("The program calculates the sum and difference of numbers.")
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

func isValidExpression(expression string) bool {
	re := regexp.MustCompile(`^[-+]?\d+(\s*[-+]\s*\d+)*`)
	return re.MatchString(expression)
}

func calculateExpression(expression string) (int, error) {
	expression = normalizeOperators(expression)
	parts := strings.Fields(expression)
	if len(parts) == 0 {
		return 0, fmt.Errorf("invalid expression")
	}

	result, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid expression")
	}

	for i := 1; i < len(parts); i += 2 {
		if i+1 >= len(parts) {
			return 0, fmt.Errorf("invalid expression")
		}
		operator := parts[i]
		num, err := strconv.Atoi(parts[i+1])
		if err != nil {
			return 0, fmt.Errorf("invalid expression")
		}
		switch operator {
		case "+":
			result += num
		case "-":
			result -= num
		default:
			return 0, fmt.Errorf("invalid expression")
		}
	}
	return result, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if strings.HasPrefix(input, "/") {
			switch input {
			case "/exit":
				fmt.Println("Bye!")
				return
			case "/help":
				showHelp()
			default:
				fmt.Println("Unknown command")
			}
			continue
		}

		if !isValidExpression(input) {
			fmt.Println("Invalid expression")
			continue
		}

		result, err := calculateExpression(input)
		if err != nil {
			fmt.Println("Invalid expression")
		} else {
			fmt.Println(result)
		}
	}
}
