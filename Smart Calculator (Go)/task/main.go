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

		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		if len(parts) == 1 {
			num, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			fmt.Println(num)
		} else if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input")
				continue
			}
			fmt.Println(num1 + num2)
		} else {
			fmt.Println("Invalid input")
		}
	}
}
