package main

import "fmt"

func main() {
	cards := createDeck()
	cards.saveToFile("cards.txt")
	c := newDeckFromFile("cards.txt")
	c.print()
	c.shuffle()
	fmt.Print("---")
	c.print()
}
