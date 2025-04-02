package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var variables = make(map[string]int)

func showHelp() {
	fmt.Println("The program calculates the sum and difference of numbers.")
	fmt.Println("It supports multiple operators (+ and -), including unary minus and repeated operators.")
	fmt.Println("You can store values in variables and use them in calculations.")
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

func isValidIdentifier(identifier string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]+$`)
	return re.MatchString(identifier)
}

func isValidAssignment(input string) bool {
	parts := strings.SplitN(input, "=", 2)
	if len(parts) != 2 {
		return false
	}
	return isValidIdentifier(strings.TrimSpace(parts[0]))
}

func evaluateExpression(expression string) (int, error) {
	expression = normalizeOperators(expression)
	parts := strings.Fields(expression)
	if len(parts) == 0 {
		return 0, fmt.Errorf("invalid expression")
	}

	var stack []int
	var lastOp rune = '+'

	for _, part := range parts {
		if num, err := strconv.Atoi(part); err == nil {
			if lastOp == '+' {
				stack = append(stack, num)
			} else {
				stack = append(stack, -num)
			}
		} else if val, exists := variables[part]; exists {
			if lastOp == '+' {
				stack = append(stack, val)
			} else {
				stack = append(stack, -val)
			}
		} else if part == "+" || part == "-" {
			lastOp = rune(part[0])
		} else {
			return 0, fmt.Errorf("invalid expression")
		}
	}

	result := 0
	for _, num := range stack {
		result += num
	}

	return result, nil
}

func handleAssignment(input string) {
	parts := strings.SplitN(input, "=", 2)
	name := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	if !isValidIdentifier(name) {
		fmt.Println("Invalid identifier")
		return
	}

	if num, err := strconv.Atoi(value); err == nil {
		variables[name] = num
	} else if val, exists := variables[value]; exists {
		variables[name] = val
	} else {
		fmt.Println("Invalid assignment")
	}
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

		if strings.Contains(input, "=") {
			if !isValidAssignment(input) {
				fmt.Println("Invalid assignment")
			} else {
				handleAssignment(input)
			}
			continue
		}

		if isValidIdentifier(input) {
			if val, exists := variables[input]; exists {
				fmt.Println(val)
			} else {
				fmt.Println("Unknown variable")
			}
			continue
		}

		result, err := evaluateExpression(input)
		if err != nil {
			fmt.Println("Invalid expression")
		} else {
			fmt.Println(result)
		}
	}
}
