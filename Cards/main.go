package main

func main() {
	// var card string = "Aces of Spades"
	// card := "Aces of Spades" //go compiler will decide what types of variable is this with ":=".
	//we can't do this we can't reassign this with ":=".
	// card := "testing"//

	//we can do this
	// card = "testing"

	// card := newCard()

	// slice example

	// card := deck{"Aces of Diamonds", newCard()}
	// card = append(card, "Testing")

	//loop example
	// for i, c := range card {

	// 	fmt.Println(i, c)
	// }

	card := newDeck()
	card.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
