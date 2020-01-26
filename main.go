package main

// func main() {
// 	// var card string = "Ace of Spades"
// 	// the line below does the same thing with different syntax, data type is infered from the right of the := / This is only ever used when Declaring a NEW variable.
// 	card := newCard()
// 	fmt.Println(card)
// }

func main() {
	cards := newDeck()
	cards.print()
}
