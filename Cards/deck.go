package main

import "fmt"

//create a new types of deck
// which is a slice of string

type deck []string

//reciver function
//print out the deck in a nice format
//the for loop will loop through the slice and print out the index and value
//of each card in the slice
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", " Heart", "Clubs", "Hearts"}
	cardValues := []string{"Aces", "One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}
