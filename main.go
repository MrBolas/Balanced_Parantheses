package main

import (
	"fmt"
	"strings"
)

const entryCharacters = "([{"
const stoppingCharacters = ")]}"

func main() {

	// slice of strings as input
	inputs := []string{
		"()()()",
		"()()",
		"",
		"(",
		")",
		"()(())",
		"()(())()()",
		"()(()))",
		"()()())",
		"()[]{}",
		"()[()()]",
		"()[(){}()]",
		"()[()({})]",
		"(){}[()()]",
		"{()[()()]}",
	}

	//iterate over slice
	for _, input := range inputs {
		// validate reverse polish notation
		valid := validate(input)
		fmt.Println(valid, " => ", input)
	}

}

func validate(input string) bool {

	//FILO First In Last out
	var stack = []string{}
	var popedValue string = ""

	for _, v := range input {
		char := string(v)

		if strings.Contains(entryCharacters, char) {
			stack = push(stack, char)
		}

		if strings.Contains(stoppingCharacters, char) {
			if len(stack) < 1 {
				return false
			}
			stack, popedValue = pop(stack)

			//validate closing bracket is the right one
			if strings.Index(entryCharacters, popedValue) != strings.Index(stoppingCharacters, char) {
				return false
			}
		}
	}

	return len(stack) == 0
}

// Push to head of stack
func push(stack []string, pushedValue string) []string {
	newStack := make([]string, 0)
	newStack = append(newStack, pushedValue)
	newStack = append(newStack, stack...)
	//fmt.Println(newStack)
	return newStack
}

// Pop from head of stack
func pop(stack []string) ([]string, string) {
	newStack := make([]string, 0)
	if len(stack) > 1 {
		newStack = append(newStack, stack[1:]...)
	}
	//fmt.Println(newStack, stack[0])
	return newStack, stack[0]
}
