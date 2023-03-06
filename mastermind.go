package main

import (
	"fmt"
	"math/rand"
)

type Dot struct {
	name string
	dot  string
}

func getGameInput(question string, number int) int {
	var response int

	fmt.Printf(question)
	fmt.Scanln(&response)

	if response == 0 {
		return number
	}

	return response
}

func getSequenceDisplayPlus(sequence []Dot) string {
	display := ""
	for _, c := range sequence {
		display = display + c.dot + " " + c.name + "  "
	}

	return display
}

func getSequenceDisplay(sequence []Dot) string {
	display := ""
	for _, c := range sequence {
		display = display + c.dot + "  "
	}

	return display
}

func createAnswer(sequenceLength int, choices []Dot) []Dot {
	var answer []Dot
	for i := 0; i < sequenceLength; i++ {
		answer = append(answer, choices[rand.Intn(len(choices))])
	}

	return answer
}

func enterColor(choices []Dot) Dot {
	var choice string

	fmt.Printf("Enter a color: ")
	fmt.Scanln(&choice)

	for _, c := range choices {
		if choice == c.name {
			return c
		}
	}

	fmt.Println("I couldn't understand your choice. Please try again!")
	fmt.Println()

	return enterColor(choices)
}

func processGuess(sequenceLength int, choices []Dot) []Dot {
	var guess []Dot
	var confirm string

	fmt.Printf("\nEnter a color as a single letter: %s\n", getSequenceDisplayPlus(choices))
	for i := 0; i < sequenceLength; i++ {
		dot := enterColor(choices)
		guess = append(guess, dot)
	}

	fmt.Printf("\nConfirm guess (y/n)? %s", getSequenceDisplay(guess))
	fmt.Scanln(&confirm)
	if confirm == "y" {
		fmt.Printf("\n")
		return guess
	}
	fmt.Printf("Guess not confirmed...trying again!\n\n")

	return processGuess(sequenceLength, choices)
}

func scoreGuess(guess []Dot, answer []Dot) []Dot {
	var score []Dot

	var answerCheck []bool
	var guessCheck []bool

	p1 := Dot{"b", "●"}
	p2 := Dot{"w", "○"}

	for n := 0; n < len(guess); n++ {
		answerCheck = append(answerCheck, true)
		guessCheck = append(guessCheck, true)
	}

	for i := 0; i < len(guess); i++ {
		if guess[i].name == answer[i].name {
			score = append(score, p1)
			answerCheck[i] = false
			guessCheck[i] = false
		}
	}

	for j := 0; j < len(guess); j++ {
		for k := 0; k < len(answer); k++ {
			if guessCheck[j] && answerCheck[k] && guess[j].name == answer[k].name {
				score = append(score, p2)
				guessCheck[j] = false
				answerCheck[k] = false
			}
		}
	}

	return score
}

func isCorrect(sequenceLength int, score []Dot) bool {
	if len(score) != sequenceLength {
		return false
	}

	for i := 0; i < sequenceLength; i++ {
		if score[i].name == "w" {
			return false
		}
	}

	return true
}

func displayGuesses(guesses [][]Dot, scores [][]Dot) {
	for i := 0; i < len(guesses); i++ {
		fmt.Printf("Guess %2d: %s Score: %s\n", i+1, getSequenceDisplay(guesses[i]), getSequenceDisplay(scores[i]))
	}
}

func main() {
	var guesses [][]Dot
	var scores [][]Dot

	var c1 = Dot{"r", "🔴"}
	var c2 = Dot{"o", "🟠"}
	var c3 = Dot{"y", "🟡"}
	var c4 = Dot{"g", "🟢"}
	var c5 = Dot{"b", "🔵"}
	var c6 = Dot{"v", "🟣"}

	choices := []Dot{c1, c2, c3, c4, c5, c6}

	rules := "Mastermind is a game where you guess a random sequence of colors."
	rules = rules + " The traditional game is a sequence of 4 with 10 guesses, but you can customize.\n"
	rules = rules + "You are scored after each guess.\n● means you have the correct color in the correct position."
	rules = rules + "\n○ means you have the correct color, but in the wrong position."
	rules = rules + "\nThe scoring pegs are sorted and don't correlate to the position of your colors.\n"

	fmt.Printf("\n%s\n", rules)

	sequenceLength := getGameInput("How long should the sequence be? (default 4)", 4)
	guessesLength := getGameInput("How many guesses should you get? (default 10)", 10)

	answer := createAnswer(sequenceLength, choices)
	fmt.Printf("\nSecret code has been generated. You have %d guesses. Good Luck!\n", guessesLength)

	for i := 0; i < guessesLength; i++ {
		guess := processGuess(sequenceLength, choices)
		guesses = append(guesses, guess)

		score := scoreGuess(guess, answer)
		scores = append(scores, score)

		displayGuesses(guesses, scores)
		if isCorrect(sequenceLength, score) {
			fmt.Printf("\n **** You won in %d guesses! ****\n\n", i+1)
			return
		}
	}
	fmt.Printf("\n **** Sorry, you lost! The answer was: %s\n\n", getSequenceDisplay(answer))
}
