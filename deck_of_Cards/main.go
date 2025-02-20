package main

func main() {
	//printDeck(createNewDeck()) // Ensure createNewDeck is accessible
	newDeck := createNewDeck()
	//printDeck(deck)
	newDeck.printDeck()
	//type conversion from [] string to deck
	// s := []string{"test"}
	// (deck(s)).printDeck()
}

// func main() {
// 	newDeck := []string{"Ace of Spaces", "Two of Diamonds"}
// 	newDeck = append(newDeck, "Three of Diamonds")
// 	fmt.Println(newDeck)
// 	printDeck(newDeck)
// }
// func printDeck(d []string) {
// 	fmt.Println(d)
// }
