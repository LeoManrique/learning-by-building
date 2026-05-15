package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

func main() {
	var lowerLimit, upperLimit int
	flag.IntVar(&lowerLimit, "lower", 1, "lower bound (inclusive)")
	flag.IntVar(&upperLimit, "upper", 100, "upper bound (inclusive)")
	flag.Parse()

	if lowerLimit > upperLimit {
		fmt.Fprintf(os.Stderr, "lower (%d) should not exceed upper (%d)\n", lowerLimit, upperLimit)
		os.Exit(2)
	}

	numberOfAttempts := 0
	notYetGuessed := true
	pickedNumber := lowerLimit + rand.IntN(upperLimit-lowerLimit+1)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Number Guessing Game!")
	fmt.Printf("I've picked a number between %d and %d, can you guess it?\n",
		lowerLimit,
		upperLimit)

	for notYetGuessed {
		if ok := scanner.Scan(); !ok {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "input error:", err)
				os.Exit(1)
			}
			fmt.Println("No more input — goodbye!")
			os.Exit(1)
		}
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("That's not a valid number, try again:")
			continue
		}

		if num > upperLimit || num < lowerLimit {
			fmt.Printf("That's not between %d and %d, try again:\n", lowerLimit, upperLimit)
			continue
		}

		numberOfAttempts++

		if num == pickedNumber {
			fmt.Printf("You guessed it after %d attempts! It was %d. See you! :)\n", numberOfAttempts, pickedNumber)
			notYetGuessed = false
		} else if num > pickedNumber {
			fmt.Println("Lower :P, try again:")
		} else {
			fmt.Println("Higher :D, try again:")
		}
	}
}
