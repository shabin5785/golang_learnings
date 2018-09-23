package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

//create a new type DECK which is a slice of string
type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, " ", card)
	}
}

//we could have written as receiver, but we
//might call deal multiple times on a deck. like a
//deal on a partial deck.. like return value of an already dealed deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)

}

func (d deck) shuffle() {

	for i := range d {
		newPos := rand.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}
