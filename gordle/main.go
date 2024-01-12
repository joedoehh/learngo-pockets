package main

import (
	"fmt"
	"os"

	"learngo-pockets/gordle/gordle"
)

const maxAttempts = 6

func main() {
	// determine corpus filename
	fileName := "corpus/english.txt"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}
	fmt.Printf("Choosing words from corpus %s\n", fileName)

	// read corpus
	corpus, err := gordle.ReadCorpus(fileName)
	if err != nil {
		panic(err)
	}

	// Create the game.
	g, err := gordle.New(os.Stdin, corpus, maxAttempts)
	if err != nil {
		panic(err)
	}

	// Run the game ! It will end when it's over.
	g.Play()
}
