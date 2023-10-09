package util

import (
	"fmt"
)

func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
