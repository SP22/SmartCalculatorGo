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
	fmt.Println("The program supports the following features:")
	fmt.Println("- Addition (+), Subtraction (-), Multiplication (*), Integer Division (/), and Power (^)")
	fmt.Println("- Support for parentheses to alter precedence")
	fmt.Println("- Repeated operators like --- or +++")
	fmt.Println("- Variable assignment and usage")
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
	return strings.Count(input, "=") == 1
}

func precedence(op string) int {
	switch op {
	case "^":
		return 3
	case "*", "/":
		return 2
	case "+", "-":
		return 1
	default:
		return 0
	}
}

func toPostfix(tokens []string) ([]string, error) {
	var output []string
	var stack []string

	for _, token := range tokens {
		switch {
		case isValidIdentifier(token) || isNumber(token):
			output = append(output, token)
		case token == "(":
			stack = append(stack, token)
		case token == ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1] != "(" {
				return nil, fmt.Errorf("Invalid expression")
			}
			stack = stack[:len(stack)-1] // pop '('
		case strings.Contains("+-*/^", token):
			for len(stack) > 0 && precedence(token) <= precedence(stack[len(stack)-1]) {
				if stack[len(stack)-1] == "(" {
					break
				}
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		default:
			return nil, fmt.Errorf("Invalid expression")
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, fmt.Errorf("Invalid expression")
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

func evalPostfix(postfix []string) (int, error) {
	var stack []int
	for _, token := range postfix {
		if isNumber(token) {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		} else if isValidIdentifier(token) {
			val, ok := variables[token]
			if !ok {
				return 0, fmt.Errorf("Unknown variable")
			}
			stack = append(stack, val)
		} else if strings.Contains("+-*/^", token) {
			if len(stack) < 2 {
				return 0, fmt.Errorf("Invalid expression")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result int
			switch token {
			case "+":
				result = left + right
			case "-":
				result = left - right
			case "*":
				result = left * right
			case "/":
				if right == 0 {
					return 0, fmt.Errorf("Division by zero")
				}
				result = left / right
			case "^":
				result = 1
				for i := 0; i < right; i++ {
					result *= left
				}
			}
			stack = append(stack, result)
		} else {
			return 0, fmt.Errorf("Invalid expression")
		}
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("Invalid expression")
	}
	return stack[0], nil
}

func isNumber(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func tokenize(expression string) ([]string, error) {
	expression = normalizeOperators(expression)
	re := regexp.MustCompile(`\d+|[a-zA-Z]+|[()+\-*/^]`)
	tokens := re.FindAllString(expression, -1)
	for _, token := range tokens {
		if strings.Count(token, "*") > 1 || strings.Count(token, "/") > 1 {
			return nil, fmt.Errorf("Invalid expression")
		}
	}
	return tokens, nil
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
	} else if isValidIdentifier(value) {
		fmt.Println("Unknown variable")
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

		tokens, err := tokenize(input)
		if err != nil {
			fmt.Println("Invalid expression")
			continue
		}
		postfix, err := toPostfix(tokens)
		if err != nil {
			fmt.Println("Invalid expression")
			continue
		}
		result, err := evalPostfix(postfix)
		if err != nil {
			fmt.Println("Invalid expression")
		} else {
			fmt.Println(result)
		}
	}
}
