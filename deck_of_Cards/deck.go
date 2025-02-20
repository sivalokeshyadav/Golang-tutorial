package main

import "fmt"

type deck []string

func createNewDeck() deck {
	newDeck := []string{"Ace of Spades", "Two of Diamonds"}
	newDeck = append(newDeck, "Three of Clubs")
	return newDeck
	//TODOassignment
	//returnanewdeckofcardsfromallsuits(Spades,Diamonds,clubs,hearts)
	//and numbers (Ace,one,two,three,four)
}

//	func printDeck(d deck) {
//		fmt.Println(d)
//	}
// func (d deck) printDeck() {
// 	//fmt.Println(d)
// 	for i, card := range d {
// 		fmt.Println(i, card)
// 	}
// }

//blank reference
func (d deck) printDeck() {
	//fmt.Println(d)
	for _, card := range d {
		fmt.Println(card)
	}
}
