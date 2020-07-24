package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/fatih/color"
)

func main() {

	var letters []string
	var hidden []string
	var indexes []int

	var guesses = 3

	fmt.Println("Welcom To GoMan Hangman Made In GO")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What Word Would You Like Them To Guess? ")
	wordToGuess, _ := reader.ReadString('\n')
	wordToGuess = strings.TrimSuffix(wordToGuess, "\r\n")
	for i := 0; i < len(wordToGuess); i++ {
		letters = append(letters, string(wordToGuess[i]))
		hidden = append(hidden, "_")
	}
	for guesses != 0 {
		color.Blue(strings.Join(hidden, " "))
		if reflect.DeepEqual(hidden, letters) {
			color.Green("You Won!")
			break
		}
		fmt.Print("Enter a Guess: ")
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSuffix(guess, "\r\n")
		for i := 0; i < len(letters); i++ {
			if letters[i] == guess {
				indexes = append(indexes, i)
			}
		}
		if len(indexes) > 0 {
			for i := 0; i < len(indexes); i++ {
				hidden[indexes[i]] = letters[indexes[i]]
			}
		} else {
			guesses--
		}
	}
	if guesses == 0 {
		fmt.Println("You ran out of guesses!")
	}
}
