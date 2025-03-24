package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
			fmt.Println("The program calculates the sum of numbers")
			continue
		}

		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		sum := 0
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			sum += num
		}
		fmt.Println(sum)
	}
}
