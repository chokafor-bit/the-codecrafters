package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}
func Title(s string) string {
	t := strings.ToLower(s)
	return strings.Title(t)
}

func reverseStr(s string) string {
	t := strings.Fields(s)
	q := ""
	r := ""
	for i := 0; i < len(t); i++ {
		r = t[i] + " "
		for j := len(r) - 1; j >= 0; j-- {
			q += string(r[j])
		}
	}
	return q
}

func snakecase(word string) string {
	words := strings.Fields(word)
	d := ""
	s := ""
	for i := 0; i < len(words); i++ {
		d += words[i] + " "
	}
	for j := 0; j < len(d)-1; j++ {
		s += string(d[j])
		s = strings.ToLower(s)
		s = strings.ReplaceAll(s, " ", "_")
	}
	return s
}
func ToLower(s string) string {
	return strings.ToLower(s)
}

func handleChoice(choice, input string, history []string) []string {
	switch choice {
	case "1":
		result := strings.ToUpper(input)
		fmt.Println("Result: " + result)
		return append(history, "UPPER  "+input+result)
	case "2":

		result := strings.ToLower(input)
		fmt.Println("Result: " + result)
		return append(history, "lower  "+input+result)
	case "3":

		result := strings.Title(input)
		fmt.Println("Result: " + result)
		return append(history, "Title  "+input+result)
	case "4":

		result := reverseStr(input)
		fmt.Println("Result: " + result)
		return append(history, "Reverse  "+input+result)
	case "5":

		result := snakecase(input)
		fmt.Println("Result: " + result)
		return append(history, "Snake "+input+result)
	case "6":

		if isPalindrom(input) {
			fmt.Println("Result: This is a palindrome!")
			return append(history, "Palindrome  "+input)
		}
		fmt.Println("Result: Not a palindrome.")
		return append(history, "Palindrome  "+input)
	case "7":

		showHistory(history)
	default:
		fmt.Println("Invalid choice, try again.")
		fmt.Println("Valid choices: >>>> 1.UPPER  2.lower  3.Title  4.Reverse  5.Snake  6.Palindrome  7.History")
	}
	return history
}

func isPalindrom(s string) bool {
	c := ""
	d := ""
	s = strings.ToLower(s)
	for _, r := range s {
		if r != ' ' {
			c += string(r)
		}
		for i := len(c) - 1; i >= 0; i-- {
			d += string(c[i])
			break
		}
	}
	return d == c
}

func showHistory(history []string) {
	fmt.Println("History:")
	if len(history) == 0 {
		fmt.Println("No Actions Are saved Yet.")
		return
	}
	for i, h := range history {
		fmt.Println(i+1, ".", h)
	}
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("SENTINEL RECOVERY UNIT ONLINE")
	var history []string
	for {
		fmt.Print("Enter text (or 'stop'): ")
		if !reader.Scan() {
			break
		}
		input := reader.Text()
		if strings.ToLower(input) == "stop" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
		}
		fmt.Println("1.UPPER 2.lower 3.Title 4.Reverse 5.Snake 6.Palindrome 7.History")
		fmt.Print("Your Choice Number: ")
		if !reader.Scan() {
			break
		}
		history = handleChoice(reader.Text(), input, history)
	}
}
