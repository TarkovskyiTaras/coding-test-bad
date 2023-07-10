package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	input := readInputFromTerminal()
	textFixed := GetNormalizedString(input)
	textFixedSlice := strings.Split(textFixed, " ")
	stringOfFirstUnique := getFirstUniqueCharsInSlice(textFixedSlice)
	firstUniqueCharInUniqueString := getFirstUniqueCharInString(stringOfFirstUnique)

	fmt.Println("********************************************")
	fmt.Printf("Output: %s\n", firstUniqueCharInUniqueString)
}

func readInputFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please copy and paste your text into the terminal. Press enter when on the new line to proceed.")

	var lines []string
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Break out of loop when \n didn't just happen to be a part of text
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			break
		}

		lines = append(lines, input)
	}

	// Concatenate the lines into a single input string
	input := strings.Join(lines, " ")

	return input
}

// GetNormalizedString gets rid of everything except letters and spaces in a given string
func GetNormalizedString(input string) string {
	text := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return r
		case r >= 'A' && r <= 'Z':
			return r
		case r == ' ':
			return r
		case r == '\n':
			return ' '
		default:
			return -1
		}
	}, input)

	re := regexp.MustCompile(`\s+`)
	textFixed := re.ReplaceAllString(text, " ")

	return textFixed
}

func getFirstUniqueCharsInSlice(words []string) string {
	var result string
	var counts map[rune]int

	for _, word := range words {
		// Reset the count for each word
		counts = make(map[rune]int)

		// Iterate over the characters in the word
		for _, char := range word {
			// Increment the count for each character
			counts[char]++
		}

		// Find the first unique character in the word
		for _, char := range word {
			if counts[char] == 1 {
				result += string(char)
				break
			}
		}
	}

	return result
}

func getFirstUniqueCharInString(word string) string {
	var result string
	counts := make(map[rune]int)

	for _, char := range word {
		counts[char]++
	}

	for _, char := range word {
		if counts[char] == 1 {
			result += string(char)
			break
		}
	}
	return result
}