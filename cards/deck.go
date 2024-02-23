package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// create  new type deck which is slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuite := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardVlaue := []string{"Ace", "Two", "Three", "Four"}
	for _, suite := range cardSuite {
		for _, value := range cardVlaue {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
func fromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), "\n"))
}
func (d deck) shuffle() {
	for i, _ := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
