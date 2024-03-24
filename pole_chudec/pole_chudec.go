package main

import (
	"fmt"
	"os"
	"strings"
)

func get_input_char() (result string) {
	var input string
	try_count := 0
	for true {
		try_count += 1
		// читаем значение в переменную
		fmt.Scanln(&input)
		if input == "END" {
			os.Exit(1)
		}
		if try_count > 50 {
			return "You have found a bug! Try again \n"
		}
		if len(input) > 1 {
			fmt.Printf("only 1 char! \n")
			continue
		}

		break
	}
	return input
}

func parse_word(word string) {
	var revealed string
	var one_char string
	fmt.Printf("next word, %d\n\n", len(word))
	chars := strings.Split(word, "")
	revealed = strings.Repeat("*", len(word))
	hidden := strings.Split(revealed, "")
	fmt.Printf("%s \n", hidden)

	for true {
		one_char = get_input_char()
		if one_char == "You have found a bug! Try again \n" {
			continue
		}
		for index, char := range chars {
			if hidden[index] == "*" {
				if one_char == char {
					hidden[index] = one_char
				}
			}
		}
		there_is_hidden_char := strings.ContainsAny(strings.Join(hidden, ""), "*")
		fmt.Printf("%s \n", hidden)
		if !there_is_hidden_char {
			break
		}
	}
	fmt.Printf("word end\n\n")

}

func main() {
	words := []string{"car", "variable", "machinelearning"}
	for _, word := range words {
		parse_word(word)
	}

}
